package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Commnad-line flag name=addr, default value=:4000 and help message
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Initialize new new servemux
	// Register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()

	// Create file server which serves files from the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// We strip the "/static" prefix from the URL path before the request reaches the file server
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Print a log message to say the server is starting
	log.Printf("Starting server on %s", *addr)

	// Start a new web server with two parameters:
	// 1. The TCP network address to listen to
	// 2. The servemux to use
	// Pass the deferenced addr pointer
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
