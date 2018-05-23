package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	http_transport "github.com/go-kit/kit/transport/http"
)

// MakeHandlerForEndpoint makes a HTTP handler for a given handler
func MakeHandlerForEndpoint(e endpoint.Endpoint) *http_transport.Server {
	return http_transport.NewServer(e,
		decodeHTTPRequest,
		encodeHTTPReply,
		http_transport.ServerErrorEncoder(httpErrorHandler),
	)
}

// decodeHTTPRequest decodes the flatbuffer flaki request.
func decodeHTTPRequest(_ context.Context, req *http.Request) (interface{}, error) {
	return req.Body, nil
}

// encodeHTTPReply encodes the flatbuffer flaki reply.
func encodeHTTPReply(_ context.Context, w http.ResponseWriter, rep interface{}) error {
	var b bytes.Buffer
	var enc = json.NewEncoder(&b)
	enc.SetIndent("", "  ")
	if err := enc.Encode(rep); err != nil {
		http.Error(w, err.Error(), 500)
		return err
	}
	fmt.Fprintf(w, "%s", b.String())
	return nil
}

// httpErrorHandler encodes the flatbuffer flaki reply when there is an error.
func httpErrorHandler(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
