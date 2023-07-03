package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

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

// logic for HTTP DELETE Method
func Delete(rw http.ResponseWriter, a any) {

	// returns the value that the interface points to
	object := reflect.Indirect(reflect.ValueOf(a))

	// check if value is nil
	if object.IsNil() {
		rw.WriteHeader(http.StatusOK) // 200 OK
		rw.Write([]byte("resource deleted successfully"))
	}

}
