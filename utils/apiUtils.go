package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// logic for POST Method
func Post[T any](rw http.ResponseWriter, req *http.Request, obj T) {

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

// logic for GET Method
func Get[T any](rw http.ResponseWriter, req *http.Request, obj T) {

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(&obj)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}
