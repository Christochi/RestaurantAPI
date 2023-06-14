package menu

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// pathnames for subroot in url endpoint
var (
	allMenuRegex = regexp.MustCompile(`^\/menu[\/]?$`) // /menu or /menu/
)

// Menu json Object
type menuJson struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Desc  string `json:"desc"` // description
}

type menu []menuJson // slice type to be used as a receiver for methods

// retuns menu object
func NewMenu() *menu {

	return new(menu)

}

// handlerfunc for menu endpoint
func (m *menu) MenuHandler(rw http.ResponseWriter, req *http.Request) {

	// inform browser to expect json
	rw.Header().Set("Content-Type", "application/json")

	// determines the function to call by the request
	switch {

	case req.Method == http.MethodPost && allMenuRegex.MatchString(req.URL.Path):
		m.PostMenu(rw, req)

	case req.Method == http.MethodGet && allMenuRegex.MatchString(req.URL.Path):
		m.GetMenu(rw, req)

	// case req.Method == http.MethodGet && specificChefRegex.MatchString(req.URL.Path):
	// 	c.GetChefByName(rw, req)

	// case req.Method == http.MethodDelete && allChefsRegex.MatchString(req.URL.Path):
	// 	c.DeleteChef(rw, req)

	// case req.Method == http.MethodDelete && specificChefRegex.MatchString(req.URL.Path):
	// 	c.DeleteChefByName(rw, req)

	default:
		m.NotFound(rw, req) // returns 501 Not Implemented
	}

}

// client send chef data using POST Method
func (m *menu) PostMenu(rw http.ResponseWriter, req *http.Request) {

	// read response body and decode json to struct
	err := json.NewDecoder(req.Body).Decode(&m)

	// error handling
	if err != nil {
		log.Fatal("error decoding into struct")
	} else {
		// server's response to client
		fmt.Fprintf(rw, "%s\n", http.StatusText(http.StatusCreated)) // 201 Created
	}

}

// sends message to client if resource does not exist or not implemented
func (m *menu) NotFound(rw http.ResponseWriter, req *http.Request) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
