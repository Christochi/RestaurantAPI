package menu

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// pathnames for subroot in url endpoint
var (
	allMenuRegex      = regexp.MustCompile(`^\/menu[\/]?$`)         // /menu or /menu/
	allBreakfastRegex = regexp.MustCompile(`^\/menu\/(breakfast)$`) // /menu/breakfast
	allLunchRegex     = regexp.MustCompile(`^\/menu\/(lunch)$`)     // /menu/lunch
	allDinnerRegex    = regexp.MustCompile(`^\/menu\/(Dinner)$`)    // /menu/dinner
	allDrinksRegex    = regexp.MustCompile(`^\/menu\/(drinks)$`)    // /menu/drinks
)

// Menu json Object
type menuJson struct {
	Type  string `json:"type"` // meal type
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

	case req.Method == http.MethodGet && allBreakfastRegex.MatchString(req.URL.Path):
		m.GetBreakfastMenu(rw, req)

	case req.Method == http.MethodGet && allLunchRegex.MatchString(req.URL.Path):
		m.GetLunchMenu(rw, req)

	case req.Method == http.MethodDelete && allMenuRegex.MatchString(req.URL.Path):
		m.DeleteMenu(rw, req)

	// case req.Method == http.MethodDelete && specificChefRegex.MatchString(req.URL.Path):
	// 	c.DeleteChefByName(rw, req)

	default:
		m.NotFound(rw, req) // returns 501 Not Implemented
	}

}

// client send menu data using POST Method
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

// client requests for menu data using GET Method
func (m *menu) GetMenu(rw http.ResponseWriter, req *http.Request) {

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(&m)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// client requests for breakfast menu
func (m *menu) GetBreakfastMenu(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := allBreakfastRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /menu/breakfast = ["/menu/breakfast", "breakfast"]
	mealType := urlSubPaths[1]

	var meal []menuJson // new slice to hold the filtered data

	for _, value := range *m {

		if strings.ToLower(value.Type) == mealType {
			meal = append(meal, value) // append to new slice
		}

	}

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(meal)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// client requests for lunch menu
func (m *menu) GetLunchMenu(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := allBreakfastRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /menu/lunch = ["/menu/lunch", "lunch"]
	mealType := urlSubPaths[1]

	var meal []menuJson // new slice to hold the filtered data

	for _, value := range *m {

		if strings.ToLower(value.Type) == mealType {
			meal = append(meal, value) // append to new slice
		}

	}

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(meal)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// client requests for dinner menu
func (m *menu) GetDinnerMenu(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := allBreakfastRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /menu/dinner = ["/menu/dinner", "dinner"]
	mealType := urlSubPaths[1]

	var meal []menuJson // new slice to hold the filtered data

	for _, value := range *m {

		if strings.ToLower(value.Type) == mealType {
			meal = append(meal, value) // append to new slice
		}

	}

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(meal)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// client requests for drinks menu
func (m *menu) GetDrinksMenu(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := allBreakfastRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /menu/drinks = ["/menu/drinks", "drinks"]
	mealType := urlSubPaths[1]

	var meal []menuJson // new slice to hold the filtered data

	for _, value := range *m {

		if strings.ToLower(value.Type) == mealType {
			meal = append(meal, value) // append to new slice
		}

	}

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(meal)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// client deletes all menu data
func (m *menu) DeleteMenu(rw http.ResponseWriter, req *http.Request) {

	// delete all element by re-initializing to nil
	*m = nil

	fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")

}

// sends message to client if resource does not exist or not implemented
func (m *menu) NotFound(rw http.ResponseWriter, req *http.Request) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
