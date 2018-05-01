package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudtrust/mock-app-back/pkg/mockback"
	"github.com/go-kit/kit/endpoint"
)

func main() {

	// Critical errors channel.
	var errc = make(chan error)
	go func() {
		var c = make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// We create the modules
	var patientModule mockback.PatientModule
	{
		patientModule = mockback.NewPatientModule()
	}

	// We create the business components
	var patientComponent mockback.PatientComponent
	{
		patientComponent = mockback.NewPatientComponent(patientModule)
	}

	// We create the endpoints
	var listAllPatientsEndpoint endpoint.Endpoint
	{
		listAllPatientsEndpoint = mockback.MakeListAllPatientsEndpoint(patientComponent)
	}

	var endpoints = mockback.Endpoints{
		ListAllPatientsEndpoint: listAllPatientsEndpoint,
	}

	// We create the HTTP server
	go func() {
		errc <- mockback.InitWeb(endpoints)
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
