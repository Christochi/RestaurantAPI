package main

import (
	"net/http"
	"log"
)

func main() {

	fileHandler := http.FileServer(http.Dir("../static"))
	http.Handle("/", noCache(http.StripPrefix("/static/", fileHandler)))

	log.Fatal(http.ListenAndServe(":3000", nil))
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