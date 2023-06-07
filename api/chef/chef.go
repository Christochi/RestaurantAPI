package chef

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Chef json Object
type ChefJson struct {
	Name  string `json:"name"`
	About string `json:"about"`
}

var Chef []ChefJson // list of chefs

// handlerfunc for chef endpoint
func ChefHandler(rw http.ResponseWriter, req *http.Request) {

	// inform browser to expect json
	rw.Header().Set("Content-Type", "application/json")

	// determines the HTTP verb
	switch req.Method {
	case "GET":
		GetMethod(rw, req)

	case "POST":
		PostMethod(rw, req)
	}

}

// client send chef data using POST Method
func PostMethod(rw http.ResponseWriter, req *http.Request) {

	// decode json to struct
	err := json.NewDecoder(req.Body).Decode(&Chef)

	// error handling
	if err != nil {
		log.Fatal("error decoding into struct")
	} else {
		// server's response to client
		fmt.Fprintf(rw, "%s\n", http.StatusText(http.StatusCreated))
	}

}

// client requests for chef data using GET Method
func GetMethod(rw http.ResponseWriter, req *http.Request) {

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(&Chef)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}
