// Utility functions for the api package

package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Christochi/error-handler/service"
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
func Create(rw http.ResponseWriter, req *http.Request, a any) error {

	// read response body and decode json to struct
	err := json.NewDecoder(req.Body).Decode(&a)
	if err != nil {
		return service.NewError(err, http.StatusUnprocessableEntity)
	}

	return nil

}

// logic for HTTP GET Method
func Get(rw http.ResponseWriter, a any) error {

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(&a)
	if err != nil {
		return service.NewError(err, http.StatusUnprocessableEntity)
	}

	return nil

}

// server's response
func ServerMessage(rw http.ResponseWriter, msg string, code int) {

	rw.WriteHeader(code) // HTTP Status Code

	_, err := rw.Write([]byte(msg)) // response body
	if err != nil {
		http.Error(rw, "Could not write data to the connection", http.StatusInternalServerError) // 500 Internal Server Error
	}

}
