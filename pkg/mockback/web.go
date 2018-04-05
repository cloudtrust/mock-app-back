package mockback

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func mock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It works!")
}

// InitWeb initializes the web service
func InitWeb() {
	var r = mux.NewRouter()
	r.HandleFunc("/", mock)
	log.Fatal(http.ListenAndServe(":8000", r))
}
