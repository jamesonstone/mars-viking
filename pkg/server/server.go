package server

import (
	"log"
	"net/http"
)

// Start starts the API server
func Start() {
	log.Println("Listening on 3000")
	http.ListenAndServe(":3000", LoadRouter())
}
