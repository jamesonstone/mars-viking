package server

import (
	"log"
	"net/http"
)

// Start starts the API server
func Start() {
	log.Println("Listening on 4000")
	http.ListenAndServe(":4000", LoadRouter())
}
