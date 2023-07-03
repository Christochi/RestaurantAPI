package webserver

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"restaurantapi/api/chef"
	"restaurantapi/api/menu"
)

// CONSTANTS
const addr string = ":3000" // default docker exposed port

var WebFilesDir string // directory of web files

// handler interface
// type apiHandler struct {
// 	c []chef.ChefJson
// 	//menu Menu
// }

// func (api *apiHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

// 	tmpl := template.Must(template.ParseFiles("static/index.html"))
// 	tmpl.Execute(rw, api)

// }

// starts server
func RunServer(dir string) {

	// stores handler that serves requests with web files
	fileHandler := http.FileServer(http.Dir(WebFilesDir))

	// returns a ServeMux for matching HTTP requests
	router := http.NewServeMux()

	// matches handler with incoming request or pattern (endpoint)
	router.Handle("/", noCache(http.StripPrefix("/", fileHandler)))
	// router.Handle("/", noCache(&apiHandler{c: chef.Chef}))
	newChef := chef.NewChef() // returns a chef object
	router.HandleFunc("/chef/", newChef.ChefHandler)
	newMenu := menu.NewMenu() // returns a menu object
	router.HandleFunc("/menu/", newMenu.MenuHandler)

	// command line flag for setting the port that the server would listen on
	listenAddr := flag.String("listenaddr", addr, "server address")
	flag.Parse() // parse the command line into the defined flag
	fmt.Println("server running on port: ", *listenAddr)

	// listens on the bind address and handles requests from incoming connections
	log.Fatal(http.ListenAndServe(*listenAddr, router))

}

// handles browser caching
func noCache(h http.Handler) http.Handler {

	// the ResponseWriter is what the server will respond with or send the client
	// Request is what the client sends to the server
	handlerFunc := func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate;")
		rw.Header().Set("pragma", "no-cache")
		rw.Header().Set("X-Content-Type-Options", "nosniff")

		h.ServeHTTP(rw, req)
	}

	return http.HandlerFunc(handlerFunc)

}
