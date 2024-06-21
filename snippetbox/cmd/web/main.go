package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Commnad-line flag name=addr, default value=:4000 and help message
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Initialize a new structured logger that writes to standard output stream and uses the default settings
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	// Print a log message to say the server is starting
	logger.Info("starting server", slog.Any("addr", *addr))

	// Start a new web server with two parameters:
	// 1. The TCP network address to listen to
	// 2. The servemux to use
	// Pass the deferenced addr pointer
	err := http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}
