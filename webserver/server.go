package webserver

import (
	"log"
	"net/http"
	"restaurantapi/api/chef"
)

// CONSTANTS
const addr string = ":3000"

var WebFilesDir string // directory of web files

// starts server
func RunServer(dir string) {

	// stores handler that serves requests with web files
	fileHandler := http.FileServer(http.Dir(WebFilesDir))

	// returns a ServeMux for matching HTTP requests
	router := http.NewServeMux()

	// matches handler with incoming request or pattern (endpoint)
	router.Handle("/", noCache(http.StripPrefix("/", fileHandler)))
	router.HandleFunc("/chef", chef.ChefHandler)

	// listens on the network address and handles requests from incoming connections
	log.Fatal(http.ListenAndServe(addr, router))

}

// handles browser caching
func noCache(h http.Handler) http.Handler {

	// the ResponseWriter is what the server will respond with or send the client
	// Request is what the client sends to the server
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate;")
		w.Header().Set("pragma", "no-cache")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handlerFunc)

}
