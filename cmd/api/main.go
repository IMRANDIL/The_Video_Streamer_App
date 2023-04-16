package main

import (
	"The_Streaming_App/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = 8080

func main() {
	r := mux.NewRouter()

	// Register a handler function for the root path ("/")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go backend")
	})

	// Define a route to stream video
	r.HandleFunc("/video/{filename}", handlers.StreamVideoHandler).Methods("GET")

	// Start the server

	log.Println("app started on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal(err)
	}
}
