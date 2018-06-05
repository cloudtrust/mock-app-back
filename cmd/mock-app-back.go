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
	"github.com/cloudtrust/mock-app-back/pkg/files"
	"github.com/cloudtrust/mock-app-back/pkg/hospitals"
	"github.com/cloudtrust/mock-app-back/pkg/patients"
	"github.com/cloudtrust/mock-app-back/pkg/shared"
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
		httpHospitals     = c.GetString("http-hospitals")
		httpDepartments   = c.GetString("http-departments")
		httpDoctors       = c.GetString("http-doctors")
		httpFiles         = c.GetString("http-files")

		// SSE
		sseEvents = c.GetString("sse-events")
	)

	// Critical errors channel.
	var errc = make(chan error)
	go func() {
		var c = make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// We create the SSE Enpoint
	var server *sse.Server
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
	var patientDatabase *patients.Database
	{
		patientDatabase, err = patients.InitDatabase(hospitalConn)
		if err != nil {
			logger.Log("error", err)
			return
		}
	}
	var hospDepDatabase *hospitals.HospDepDatabase
	{
		hospDepDatabase, err = hospitals.InitHospDepDatabase(hospitalConn)
		if err != nil {
			logger.Log("error", err)
			return
		}
	}
	var doctorsDatabase *hospitals.DoctorsDatabase
	{
		doctorsDatabase, err = hospitals.InitDoctorsDatabase(hospitalConn)
		if err != nil {
			logger.Log("error", err)
			return
		}
	}
	var filesDatabase *files.Database
	{
		filesDatabase, err = files.InitDatabase(medifilesConn)
		if err != nil {
			logger.Log("error", err)
			return
		}
	}

	// We create the modules.
	var patientsModule patients.Module
	{
		patientsModule = patients.NewModule(*patientDatabase)
	}
	var hospDepModule hospitals.HospDepModule
	{
		hospDepModule = hospitals.NewHospDepModule(*hospDepDatabase)
	}
	var doctorsModule hospitals.DoctorsModule
	{
		doctorsModule = hospitals.NewDoctorModule(*doctorsDatabase)
	}
	var filesModule files.Module
	{
		filesModule = files.NewModule(*filesDatabase)
	}

	// We create the business components
	var patientsComponent patients.Component
	{
		patientsComponent = patients.NewComponent(patientsModule)
	}
	var hospitalsComponent hospitals.Component
	{
		hospitalsComponent = hospitals.NewComponent(hospDepModule, doctorsModule)
	}
	var filesComponent files.Component
	{
		filesComponent = files.NewComponent(filesModule)
	}

	// We create the endpoints
	var listAllPatientsEndpoint endpoint.Endpoint
	{
		listAllPatientsEndpoint = patients.MakeListAllEndpoint(patientsComponent)
	}
	var listAllHospitalsEndpoint endpoint.Endpoint
	{
		listAllHospitalsEndpoint = hospitals.MakeListAllHospitalsEndpoint(hospitalsComponent)
	}
	var listAllDepartmentsEndpoint endpoint.Endpoint
	{
		listAllDepartmentsEndpoint = hospitals.MakeListAllDepartmentsEndpoint(hospitalsComponent)
	}
	var listAllDoctorsEndpoint endpoint.Endpoint
	{
		listAllDoctorsEndpoint = hospitals.MakeListAllDoctorsEndpoint(hospitalsComponent)
	}
	var listAllFilesEndpoint endpoint.Endpoint
	{
		listAllFilesEndpoint = files.MakeListAllEndpoint(filesComponent)
	}

	// We create the HTTP server
	go func() {
		logger.Log("msg", "Starting web server...")
		var r = mux.NewRouter()

		// We handle the endpoints
		r.Handle(httpPatients, shared.MakeHandlerForEndpoint(listAllPatientsEndpoint))
		r.Handle(httpHospitals, shared.MakeHandlerForEndpoint(listAllHospitalsEndpoint))
		r.Handle(httpDepartments, shared.MakeHandlerForEndpoint(listAllDepartmentsEndpoint))
		r.Handle(httpDoctors, shared.MakeHandlerForEndpoint(listAllDoctorsEndpoint))
		r.Handle(httpFiles, shared.MakeHandlerForEndpoint(listAllFilesEndpoint))

		// We let the front-end access the back-end
		var c = cors.New(cors.Options{
			AllowedOrigins:   []string{httpAllowedOrigin},
			AllowCredentials: true,
			Debug:            true,
		})
		var h = c.Handler(r)

		errc <- http.ListenAndServe(httpAddr, h)
	}()

	// Sends random messages in SSE channel 1. To be removed later on.
	go func() {
		rand.Seed(time.Now().UTC().UnixNano())
		for {
			time.Sleep(time.Duration(5) * time.Second)
			shared.SendMessage(server, logger, sseEvents, 1, fmt.Sprintf("Ping %d!", rand.Intn(9999)))
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
	v.SetDefault("http-hospitals", "/hospitals")
	v.SetDefault("http-departments", "/departments")
	v.SetDefault("http-doctors", "/doctors")
	v.SetDefault("http-files", "/files")

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
