package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// logic for HTTP POST Method
// any is an interface for any type
func Post(rw http.ResponseWriter, req *http.Request, a any) {

	// read response body and decode json to struct
	err := json.NewDecoder(req.Body).Decode(&a)

	// error handling
	if err != nil {
		log.Fatal("error decoding into struct")
	} else {
		// server's response to client
		fmt.Fprintf(rw, "%s\n", http.StatusText(http.StatusCreated)) // 201 Created
	}

}

// logic for HTTP GET Method
func Get(rw http.ResponseWriter, req *http.Request, a any) {

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(&a)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// logic for HTTP DELETE Method
func Delete(rw http.ResponseWriter, req *http.Request, a any) {

	fmt.Fprintf(rw, "%v\n", &a)
	fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")

}

// sends status message to client if resource does not exist or not implemented
func NotImplemented(rw http.ResponseWriter, req *http.Request, a any) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
