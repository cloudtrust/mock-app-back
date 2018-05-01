package mockback

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/rs/cors"
)

var server *sse.Server

// InitSseEndpoint initializes the SSE Endpoint
func InitSseEndpoint() error {

	log.Print("Starting SSE Endpoint...")

	// Create the server.
	server = sse.NewServer(nil)
	defer server.Shutdown()

	// Register with /events endpoint.
	http.Handle("/events/", server)

	// We allow the front-end to access the endpoint
	var c = cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true,
	})
	var h = c.Handler(server)

	return http.ListenAndServe(":3000", h)
}

// SendMessage sends a message "message" on channel "channel"
func SendMessage(channel uint, message string) {
	if server != nil {
		server.SendMessage(fmt.Sprintf("/events/channel-%d", channel), sse.SimpleMessage(message))
		log.Printf("Sent message \"%s\" to channel %d.", message, channel)
	} else {
		log.Fatal("Initialize the SSE endpoint before trying to send messages!")
	}
}
