package chef

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// pathnames for subroot in url endpoint
var (
	allChefsRegex     = regexp.MustCompile(`^\/chef[\/]?$`)         // /chef or /chef/
	specificChefRegex = regexp.MustCompile(`^\/chef\/([A-Za-z]+)$`) // /chef/job, /chef/Cynthia
)

// Chef json Object
type chefJson struct {
	Name  string `json:"name"`
	About string `json:"about"`
}

type chef []chefJson // new slice type to be used as a receiver for methods

// retuns slice of chef object
func NewChef() *chef {

	return new(chef)

}

// handlerfunc for chef endpoint
func (c *chef) ChefHandler(rw http.ResponseWriter, req *http.Request) {

	// inform browser to expect json
	rw.Header().Set("Content-Type", "application/json")

	// determines the function to call by the request
	switch {

	case req.Method == http.MethodGet && allChefsRegex.MatchString(req.URL.Path):
		c.GetChef(rw, req)

	case req.Method == http.MethodGet && specificChefRegex.MatchString(req.URL.Path):
		c.GetChefByName(rw, req)

	case req.Method == http.MethodPost && allChefsRegex.MatchString(req.URL.Path):
		c.PostChef(rw, req)

	case req.Method == http.MethodDelete && allChefsRegex.MatchString(req.URL.Path):
		c.DeleteChef(rw, req)

	case req.Method == http.MethodDelete && specificChefRegex.MatchString(req.URL.Path):
		c.DeleteChefByName(rw, req)

	default:
		c.notFound(rw, req) // returns 501 Not Implemented
	}

}

// client send chef data using POST Method
func (c *chef) PostChef(rw http.ResponseWriter, req *http.Request) {

	// read response body and decode json to struct
	err := json.NewDecoder(req.Body).Decode(&c)

	// error handling
	if err != nil {
		log.Fatal("error decoding into struct")
	} else {
		// server's response to client
		fmt.Fprintf(rw, "%s\n", http.StatusText(http.StatusCreated)) // 201 Created
	}

}

// client requests for chef data using GET Method
func (c *chef) GetChef(rw http.ResponseWriter, req *http.Request) {

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(&c)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// client requests for specific chef
func (c *chef) GetChefByName(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := specificChefRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /user/job = ["/user/job", "job"]
	name := urlSubPaths[1]

	var chefNames []chefJson // new slice to hold the filtered data

	for _, value := range *c {

		if value.Name == name {
			chefNames = append(chefNames, value) // append to new slice
		}

	}

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(chefNames)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// client deletes all chef data
func (c *chef) DeleteChef(rw http.ResponseWriter, req *http.Request) {

	// delete all element by re-initializing to nil
	*c = nil

	fmt.Fprintln(rw, "Data is deleted", "\n", *c)

}

// client deletes a specific chef
func (c *chef) DeleteChefByName(rw http.ResponseWriter, req *http.Request) {}

// sends message to client if request does not exist or not implemented
func (c *chef) notFound(rw http.ResponseWriter, req *http.Request) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
