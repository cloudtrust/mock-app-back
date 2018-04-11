package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cloudtrust/mock-app-back/pkg/mockback"
)

func main() {
	go mockback.InitSseEndpoint()
	go func() {
		rand.Seed(42)
		for {
			time.Sleep(time.Duration(5) * time.Second)
			mockback.SendMessage(1, fmt.Sprintf("Ping %d!", rand.Intn(9999)))
		}
	}()

	mockback.InitWeb()
}
