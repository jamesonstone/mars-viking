package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// LoadRouter loads the routes and returns the router
func LoadRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/testjson", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body",
					http.StatusInternalServerError)
			}
			fmt.Println(string(body))
			err = ioutil.WriteFile("./pkg/config/test.json", body, 0644)
			if err != nil {
				http.Error(w, "Error reading request body",
					http.StatusInternalServerError)
			}

		} else {
			http.ServeFile(w, r, "./pkg/config/test.json")
		}
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	return r
}
