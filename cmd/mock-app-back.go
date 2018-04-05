package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func mock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It works!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mock)
	log.Fatal(http.ListenAndServe(":8000", r))
}
