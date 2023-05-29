package webserver

import (
	"net/http"
	"log"
)

// CONSTANTS
const addr string = ":3000"

var WebFilesDir string // directory of web files
var FileHandler http.Handler // placeholder for web files

// starts server
func RunServer(dir string) {

	// stores handler that serves requests with web files
	FileHandler = http.FileServer(http.Dir(WebFilesDir))

	// matches handler with incoming request (endpoint)
	http.Handle("/", noCache(http.StripPrefix("/", FileHandler)))

	// listens on the network address and handles requests from incoming connections
	log.Fatal(http.ListenAndServe(addr, nil))	

}

// handles browser caching
func noCache(h http.Handler) http.Handler {

	// the ResponseWriter is what the server will respond with or send the client
	// Request is what the client sends to the server
	handlerFunc := func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate;")
		w.Header().Set("pragma", "no-cache")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		h.ServeHTTP(w,r)
	}

	return http.HandlerFunc(handlerFunc)

}