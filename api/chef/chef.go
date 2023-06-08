package chef

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var (
	allChefsRegex     = regexp.MustCompile(`^\/chef\/$`)          // /chef/
	specificChefRegex = regexp.MustCompile(`^\/chef\/[A-Za-z]+$`) // /chef/job, /chef/Cynthia
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

	// determines the function to call by the request
	switch {

	case req.Method == http.MethodGet && allChefsRegex.MatchString(req.URL.Path):
		GetChef(rw, req)

	case req.Method == http.MethodGet && specificChefRegex.MatchString(req.URL.Path):
		GetChefByName(rw, req)

	case req.Method == http.MethodPost && allChefsRegex.MatchString(req.URL.Path):
		PostChef(rw, req)

	case req.Method == http.MethodDelete && allChefsRegex.MatchString(req.URL.Path):
		DeleteChef(rw, req)

	case req.Method == http.MethodDelete && specificChefRegex.MatchString(req.URL.Path):
		DeleteChefByName(rw, req)

	default:
		notFound(rw, req) // returns 501 Not Implemented
	}

}

// client send chef data using POST Method
func PostChef(rw http.ResponseWriter, req *http.Request) {

	// decode json to struct
	err := json.NewDecoder(req.Body).Decode(&Chef)

	// error handling
	if err != nil {
		log.Fatal("error decoding into struct")
	} else {
		// server's response to client
		fmt.Fprintf(rw, "%s\n", http.StatusText(http.StatusCreated)) // Created
	}

}

// client requests for chef data using GET Method
func GetChef(rw http.ResponseWriter, req *http.Request) {

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(&Chef)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// client requests for specific chef
func GetChefByName(rw http.ResponseWriter, req *http.Request) {}

// client deletes all chef data
func DeleteChef(rw http.ResponseWriter, req *http.Request) {}

// client deletes a specific chef
func DeleteChefByName(rw http.ResponseWriter, req *http.Request) {}

// sends message to client if request does not exist or not implemented
func notFound(rw http.ResponseWriter, req *http.Request) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
