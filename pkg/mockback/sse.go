package mockback

import (
	"fmt"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-kit/kit/log"
)

// SendMessage sends a message "message" on channel "channel"
func SendMessage(server *sse.Server, logger log.Logger, sseEvents string, channel uint, message string) {
	if server != nil {
		server.SendMessage(fmt.Sprintf("%s/channel-%d", sseEvents, channel), sse.SimpleMessage(message))
		logger.Log("msg", fmt.Sprintf("Sent message \"%s\" to channel %d.", message, channel))
	} else {
		logger.Log("error", "Initialize the SSE endpoint before trying to send messages!")
	}
}
