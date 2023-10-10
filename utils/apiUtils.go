// Utility functions for the api package

package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Logger Struct
type requestLogger struct {
	info *log.Logger
}

// return info field
func InfoLog() *log.Logger {

	// instantiate logger object
	requestLogger := requestLogger{info: log.New(os.Stdout, "REQUEST INFO: ", log.Ltime)}
	return requestLogger.info

}

// logic for HTTP POST Method
// any is an interface for any type
func Post(rw http.ResponseWriter, req *http.Request, a any) {

	// read response body and decode json to struct
	err := json.NewDecoder(req.Body).Decode(&a)

	// error handling
	if err != nil {
		http.Error(rw, "error decoding into struct", http.StatusUnprocessableEntity) // 422 Unprocessable Entity
	} else {
		// server's response to client
		fmt.Fprintf(rw, "%s\n", http.StatusText(http.StatusCreated)) // 201 Created
	}

}

// logic for HTTP GET Method
func Get(rw http.ResponseWriter, a any) {

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(&a)

	// error handling
	if err != nil {
		http.Error(rw, "error encoding into json", http.StatusUnprocessableEntity) // 422 Unprocessable Entity
	}

}

// server's response
func ServerMessage(rw http.ResponseWriter, msg string, code int) {

	rw.WriteHeader(code) // HTTP Status Code

	_, err := rw.Write([]byte(msg)) // response body
	if err != nil {
		http.Error(rw, "Could not write data to the connection", http.StatusInternalServerError) // 500 Internal Server Error
	}

}
