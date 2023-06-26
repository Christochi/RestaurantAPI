package utils

import (
	"encoding/json"
	"fmt"
	"log"
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

	// returns the value that the interface points to
	object := reflect.Indirect(reflect.ValueOf(a))

	// check if value is nil
	if object.IsNil() {
		fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")
	}

}

func DeleteItem(rw http.ResponseWriter, req *http.Request, a any, str string) {

	// returns the value that the interface points to
	object := reflect.Indirect(reflect.ValueOf(a))
	k := object.Kind()

	if k == reflect.Slice {
		// obj := object.MapRange()
		fmt.Fprintf(rw, "%v", true)
	}

	// for index, value := range object {

	// 	// remove whitespaces and returns lower case of the string
	// 	if strings.ToLower(strings.ReplaceAll(value.Meal, " ", "")) == str {
	// 		// delete an element
	// 		(*m)[index] = (*m)[len(*m)-1] // replace the element with the last element
	// 		*m = (*m)[:len(*m)-1]         // reinitialize the array with all the elements excluding last element

	// 		fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")

	// 		return // exit function call
	// 	}

	// }

	// rw.WriteHeader(http.StatusNotFound)                    // 404
	// rw.Write([]byte(http.StatusText(http.StatusNotFound))) // NotFound

}

// sends status message to client if resource does not exist or not implemented
func NotImplemented(rw http.ResponseWriter, req *http.Request, a any) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
