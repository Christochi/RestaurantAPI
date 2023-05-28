package webserver

import (
	"net/http"
	"log"
)

// CONSTANTS
const addr string = ":3000"

// starts server
func RunServer() {

	// the web files are processed relative to the dir of the caller (main.go) 
	fileHandler := http.FileServer(http.Dir("./static"))

	// matches handler with incoming request (endpoint)
	http.Handle("/", noCache(http.StripPrefix("/", fileHandler)))

	// listens on the network address and handles requests from incoming connections
	log.Fatal(http.ListenAndServe(addr, nil))	

}

func noCache(h http.Handler) http.Handler {

	// the ResponseWriter is what the server will respond with or send the client
	// Resquest is what the client sends to the server
	handlerFunc := func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate;")
		w.Header().Set("pragma", "no-cache")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		h.ServeHTTP(w,r)
	}

	return http.HandlerFunc(handlerFunc)

}