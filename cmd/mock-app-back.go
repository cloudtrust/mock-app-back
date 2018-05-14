package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	sse "github.com/alexandrevicenzi/go-sse"
	"github.com/cloudtrust/mock-app-back/pkg/mockback"
	"github.com/cloudtrust/mock-app-back/pkg/patients"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

	// Logger.
	var logger = log.NewJSONLogger(os.Stdout)
	{
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}
	defer logger.Log("msg", "goodbye")

	// Configurations.
	var c = config(log.With(logger, "unit", "config"))
	var (
		// Component
		httpAddr = c.GetString("component-http-host-port")
		sseAddr  = c.GetString("component-sse-host-port")

		// Cockroach
		cockroachHostPort    = c.GetString("cockroach-host-port")
		cockroachUsername    = c.GetString("cockroach-username")
		cockroachPassword    = c.GetString("cockroach-password")
		cockroachHospitalDB  = c.GetString("cockroach-hospital-database")
		cockroachMedifilesDB = c.GetString("cockroach-medifiles-database")

		// HTTP
		httpAllowedOrigin = c.GetString("http-allowed-origin")
		httpPatients      = c.GetString("http-patients")

		// SSE
		sseEvents = c.GetString("sse-events")
	)

	_ = cockroachMedifilesDB

	// Critical errors channel.
	var errc = make(chan error)
	go func() {
		var c = make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	type Cockroach interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
		QueryRow(query string, args ...interface{}) *sql.Row
		Query(query string, args ...interface{}) (*sql.Rows, error)
	}

	var err error

	// We establish the connetion to the hospital db.
	var hospitalConn Cockroach
	logger.Log("msg", "Connecting to hospital database...")
	hospitalConn, err = sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", cockroachUsername, cockroachPassword, cockroachHostPort, cockroachHospitalDB))
	if err != nil {
		logger.Log("error", err)
		return
	}
	logger.Log("msg", "Connected to hospital database!")

	// We establish the connection to the medifiles db.
	var medifilesConn Cockroach
	logger.Log("msg", "Connecting to medifiles database...")
	medifilesConn, err = sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", cockroachUsername, cockroachPassword, cockroachHostPort, cockroachMedifilesDB))
	if err != nil {
		logger.Log("error", err)
		return
	}
	logger.Log("msg", "Connected to medifiles database!")
	_ = medifilesConn

	// We create the database modules.
	var patientDatabase *patients.CockroachModule
	{
		patientDatabase, err = patients.InitDatabase(hospitalConn)
		if err != nil {
			logger.Log("error", err)
			return
		}
	}

	// We create the modules.
	var patientModule patients.Module
	{
		patientModule = patients.NewModule(*patientDatabase)
	}

	// We create the business components
	var patientComponent patients.Component
	{
		patientComponent = patients.NewComponent(patientModule)
	}

	// We create the endpoints
	var listAllPatientsEndpoint endpoint.Endpoint
	{
		listAllPatientsEndpoint = patients.MakeListAllPatientsEndpoint(patientComponent)
	}

	// We create the HTTP server
	go func() {
		logger.Log("msg", "Starting web server...")
		var r = mux.NewRouter()

		// We handle the endpoints
		var listAllPatientsHandler http.Handler
		{
			listAllPatientsHandler = patients.MakeListAllPatientsHandler(listAllPatientsEndpoint)
		}
		r.Handle(httpPatients, listAllPatientsHandler)

		// We let the front-end access the back-end
		var c = cors.New(cors.Options{
			AllowedOrigins:   []string{httpAllowedOrigin},
			AllowCredentials: true,
			Debug:            true,
		})
		var h = c.Handler(r)

		errc <- http.ListenAndServe(httpAddr, h)
	}()

	var server *sse.Server

	// We create the SSE Enpoint (WIP)
	// Note : No need to review this : This is just a PoC that will be turned into something useful eventually.
	go func() {
		logger.Log("msg", "Starting SSE Endpoint...")

		// Create the server.
		server = sse.NewServer(nil)
		defer server.Shutdown()

		// Register with /events endpoint.
		http.Handle(sseEvents, server)

		// We allow the front-end to access the endpoint
		var c = cors.New(cors.Options{
			AllowedOrigins:   []string{httpAllowedOrigin},
			AllowCredentials: true,
			Debug:            true,
		})
		var h = c.Handler(server)

		errc <- http.ListenAndServe(sseAddr, h)
	}()
	go func() {
		rand.Seed(time.Now().UTC().UnixNano())
		for {
			time.Sleep(time.Duration(5) * time.Second)
			mockback.SendMessage(server, logger, sseEvents, 1, fmt.Sprintf("Ping %d!", rand.Intn(9999)))
		}
	}()

	logger.Log("error", <-errc)
}

func config(logger log.Logger) *viper.Viper {
	logger.Log("msg", "load configuration and command args")

	var v = viper.New()

	// Component default.
	v.SetDefault("config-file", "./configs/mock-app-back.yml")
	v.SetDefault("component-http-host-port", "0.0.0.0:8000")
	v.SetDefault("component-sse-host-port", "0.0.0.0:3000")

	// Cockroach.
	v.SetDefault("cockroach", false)
	v.SetDefault("cockroach-host-port", "")
	v.SetDefault("cockroach-username", "")
	v.SetDefault("cockroach-password", "")
	v.SetDefault("cockroach-hospital-database", "")
	v.SetDefault("cockroach-medifiles-database", "")

	// HTTP.
	v.SetDefault("http-allowed-origin", "http://localhost:4200")
	v.SetDefault("http-patients", "/patients")

	// SSE.
	v.SetDefault("sse-events", "/events")

	// First level of override.
	pflag.String("config-file", v.GetString("config-file"), "The configuration file path can be relative or absolute.")
	v.BindPFlag("config-file", pflag.Lookup("config-file"))
	pflag.Parse()

	// Load config.
	v.SetConfigFile(v.GetString("config-file"))
	var err = v.ReadInConfig()
	if err != nil {
		logger.Log("error", err)
	}

	// If the host/port is not set, we consider the components deactivated.
	v.Set("cockroach", v.GetString("cockroach-host-port") != "")

	// Log config in alphabetical order.
	var keys = v.AllKeys()
	sort.Strings(keys)

	for _, k := range keys {
		logger.Log(k, v.Get(k))
	}

	return v
}
