package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Define a home handler
// Writes byte slice containing "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, _ *http.Request) {
	// Can only be called once per request
	// Once written, it cannot be modified
	// Must be before writing the response
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet..."))
}
