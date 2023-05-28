package webserver

import (
	"net/http"
	"log"
)

// CONSTANTS
const addr string = ":3000"

// the web files are processed relative to the dir of the caller 
var WebFilesDir string = "./static"
var FileHandler = http.FileServer(http.Dir(WebFilesDir))


// starts server
func RunServer() {

	// matches handler with incoming request (endpoint)
	http.Handle("/", noCache(http.StripPrefix("/", fileHandler)))

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