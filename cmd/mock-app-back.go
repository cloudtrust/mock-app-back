package main

import (
	"time"

	"github.com/cloudtrust/mock-app-back/pkg/mockback"
)

func main() {
	go mockback.InitSseEndpoint()
	go func() {
		for {
			time.Sleep(time.Duration(5) * time.Second)
			mockback.SendMessage(1, "Ping!")
		}
	}()

	mockback.InitWeb()
}
