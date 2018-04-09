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

// Sse returns Server-Send events every minute. For now, it's just a dummy answer since I didn't find a good SSE library compatible with Gorilla Mux
func Sse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	fmt.Fprintf(w, "A message")
}

// InitWeb initializes the web service
func InitWeb() {
	var r = mux.NewRouter()
	r.HandleFunc("/", Root)
	r.HandleFunc("/hospitals", ListHospitals)
	r.HandleFunc("/events/channel-1", Sse)
	var c = cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true,
	})
	var h = c.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", h))
}
