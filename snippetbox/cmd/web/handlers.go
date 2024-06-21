package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a home handler
// Writes byte slice containing "Hello from Snippetbox" as the response body
// Change the signature of home controller so it is defined as a method against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Use Execute() method to write template content as the response body
	// Last parameter is data to pass to the template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, _ *http.Request) {
	// Can only be called once per request
	// Once written, it cannot be modified
	// Must be before writing the response
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet..."))
}
