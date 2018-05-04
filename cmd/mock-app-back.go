package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudtrust/mock-app-back/pkg/mockback"
	"github.com/cloudtrust/mock-app-back/pkg/patients"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {

	// Critical errors channel.
	var errc = make(chan error)
	go func() {
		var c = make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// We establish the cockroach connection
	type Cockroach interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
		QueryRow(query string, args ...interface{}) *sql.Row
	}
	var cockroachConn Cockroach
	var err error
	log.Print("Connecting to database...")
	cockroachConn, err = sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", "root", "", "localhost:26257", "mockappdb"))
	if err != nil {
		log.Fatal(err)
		return
	} else {
		log.Print("Connected!")
	}

	// We create the modules
	var patientModule patients.Module
	{
		patientModule = patients.NewModule(cockroachConn)
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
		log.Print("Starting web server...")
		var r = mux.NewRouter()

		// We handle the endpoints
		var listAllPatientsHandler http.Handler
		{
			listAllPatientsHandler = patients.MakeListAllPatientsHandler(listAllPatientsEndpoint)
		}
		r.Handle("/patients", listAllPatientsHandler)

		// We let the front-end access the back-end
		var c = cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:4200"},
			AllowCredentials: true,
			Debug:            true,
		})
		var h = c.Handler(r)

		errc <- http.ListenAndServe(":8000", h)
	}()

	// We create the SSE Enpoint (wip)
	go func() {
		errc <- mockback.InitSseEndpoint()
	}()
	go func() {
		rand.Seed(42)
		for {
			time.Sleep(time.Duration(5) * time.Second)
			mockback.SendMessage(1, fmt.Sprintf("Ping %d!", rand.Intn(9999)))
		}
	}()

	log.Fatal(<-errc)
}
