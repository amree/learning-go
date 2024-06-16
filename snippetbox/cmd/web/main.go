package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize new new servemux
	// Register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Print a log message to say the server is starting
	log.Print("Starting server on :4000")

	// Start a new web server with two parameters:
	// 1. The TCP network address to listen to
	// 2. The servemux to use
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
