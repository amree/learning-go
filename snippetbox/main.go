package main

import (
	"log"
	"net/http"
)

// Define a home handler
// Writes byte slice containing "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Initialize new new servemux
	// Register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Print a log message to say the server is starting
	log.Print("Starting server on :4000")

	// Start a new web server with two parameters:
	// 1. The TCP network address to listen to
	// 2. The servemux to use
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
