package mockback

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Mock returns a dummy message to confirm the web server works
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It works!")
}

// InitWeb initializes the web service
func InitWeb() {
	var r = mux.NewRouter()
	r.HandleFunc("/", Root)
	log.Fatal(http.ListenAndServe(":8000", r))
}
