package mockback

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

// InitWeb initializes the web server
func InitWeb() error {
	log.Print("Starting web server...")
	var r = mux.NewRouter()
	r.HandleFunc("/", Root)
	r.HandleFunc("/hospitals", ListHospitals)
	var c = cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true,
	})
	var h = c.Handler(r)
	return http.ListenAndServe(":8000", h)
}
