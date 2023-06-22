package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Object interface{}

// logic for HTTP POST Method
func Post[T any](rw http.ResponseWriter, req *http.Request, obj *T) {

	// read response body and decode json to struct
	err := json.NewDecoder(req.Body).Decode(&obj)

	// error handling
	if err != nil {
		log.Fatal("error decoding into struct")
	} else {
		// server's response to client
		fmt.Fprintf(rw, "%s\n", http.StatusText(http.StatusCreated)) // 201 Created
	}

}

// logic for HTTP GET Method
func Get[T any](rw http.ResponseWriter, req *http.Request, obj *T) {

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(&obj)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// logic for HTTP DELETE Method
func Delete[T any](rw http.ResponseWriter, req *http.Request, obj *T) {

	fmt.Fprintf(rw, "%v\n", *obj)
	fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")

}

// sends status message to client if resource does not exist or not implemented
func NotFound[T any](rw http.ResponseWriter, req *http.Request, obj T) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
