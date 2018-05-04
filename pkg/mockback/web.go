package mockback

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	http_transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Root returns a dummy message to confirm the web server works
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It works!")
}

// ListHospitals returns all the hospitals in JSON form
func ListHospitals(w http.ResponseWriter, r *http.Request) {
	var b bytes.Buffer
	var enc = json.NewEncoder(&b)
	enc.SetIndent("", "  ")
	if err := enc.Encode(GetDummyData()); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		fmt.Fprintf(w, "%s", b.String())
	}
}

// MakeListAllPatientsHandler makes a HTTP handler for the ListAllPatients endpoint.
func MakeListAllPatientsHandler(e endpoint.Endpoint) *http_transport.Server {
	return http_transport.NewServer(e,
		decodeHTTPRequest,
		encodeHTTPReply,
		http_transport.ServerErrorEncoder(httpErrorHandler),
	)
}

// InitWeb initializes the web server
func InitWeb(endpoints Endpoints) error {
	log.Print("Starting web server...")
	var r = mux.NewRouter()
	r.HandleFunc("/", Root)
	r.HandleFunc("/hospitals", ListHospitals)

	var listAllPatientsHandler http.Handler
	{
		listAllPatientsHandler = MakeListAllPatientsHandler(endpoints.ListAllPatientsEndpoint)
	}
	r.Handle("/patients", listAllPatientsHandler)

	var c = cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true,
	})
	var h = c.Handler(r)
	return http.ListenAndServe(":8000", h)
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
