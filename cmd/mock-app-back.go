package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/signal"
	"syscall"
	"log"

	"github.com/cloudtrust/mock-app-back/pkg/mockback"
)

func main() {

	// Critical errors channel.
	var errc = make(chan error)
	go func() {
		var c = make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

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

	go func() {
		errc <- mockback.InitWeb()
	}()

	log.Fatal(<-errc)
}
