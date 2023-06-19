package menu

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"restaurantapi/utils"
	"strings"
)

// pathnames for subroot in url endpoint
var (
	allMenuRegex = regexp.MustCompile(`^\/menu[\/]?$`) // /menu or /menu/
	mealType     = regexp.MustCompile(`^\/menu\/([A-Za-z]+)$`)
	mealRegex    = regexp.MustCompile(`^\/menu\/(breakfast|lunch|drinks|dinner)\/([A-Za-z]+)$`) // /menu/<anymealtype>/burger
)

// Menu json Object
type menuJson struct {
	Type  string `json:"type"` // meal type
	Meal  string `json:"meal"`
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

	case req.Method == http.MethodGet && mealType.MatchString(req.URL.Path):
		m.GetMealType(rw, req)

	case req.Method == http.MethodDelete && allMenuRegex.MatchString(req.URL.Path):
		m.DeleteMenu(rw, req)

	case req.Method == http.MethodDelete && mealRegex.MatchString(req.URL.Path):
		m.DeleteMeal(rw, req)

	default:
		m.NotFound(rw, req) // returns 501 Not Implemented
	}

}

// client send menu data using POST Method
func (m *menu) PostMenu(rw http.ResponseWriter, req *http.Request) {

	// read and decode to struct
	utils.Post[*menu](rw, req, m)

}

// client requests for menu data using GET Method
func (m *menu) GetMenu(rw http.ResponseWriter, req *http.Request) {

	// encode to json and rw sends the json
	utils.Get[*menu](rw, req, m)

}

// gets the list of menu for a meal type
func (m *menu) GetMealType(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := mealType.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /menu/breakfast = ["/menu/breakfast", "breakfast"]
	mealTypeName := urlSubPaths[1]

	var meal []menuJson // new slice to hold the filtered data

	for _, value := range *m {

		if strings.ToLower(value.Type) == mealTypeName {
			meal = append(meal, value) // append to new slice

			// encode to json and rw sends the json
			err := json.NewEncoder(rw).Encode(meal)

			// error handling
			if err != nil {
				log.Fatal("error encoding into json")
			}

			return
		}

	}

	rw.WriteHeader(http.StatusNotFound)                    // 404
	rw.Write([]byte(http.StatusText(http.StatusNotFound))) // NotFound

}

// client deletes all menu data
func (m *menu) DeleteMenu(rw http.ResponseWriter, req *http.Request) {

	// delete all element by re-initializing to nil
	*m = nil

	if *m != nil {
		fmt.Fprintf(rw, "%s\n", http.StatusText(http.StatusInternalServerError)) // 500 Internal Server Error
	} else {
		fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")
	}

}

func (m *menu) DeleteMeal(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := mealRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the third index
	// example: /menu/breakfast = ["/menu/lunch/burger", "lunch", "burger"]
	meal := urlSubPaths[2]

	for index, value := range *m {

		// remove whitespaces and returns lower case of the string
		if strings.ToLower(strings.ReplaceAll(value.Meal, " ", "")) == meal {
			// delete an element
			(*m)[index] = (*m)[len(*m)-1] // replace the element with the last element
			*m = (*m)[:len(*m)-1]         // reinitialize the array with all the elements excluding last element

			fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")

			return
		}

	}

	rw.WriteHeader(http.StatusNotFound)                    // 404
	rw.Write([]byte(http.StatusText(http.StatusNotFound))) // NotFound

}

// sends message to client if resource does not exist or not implemented
func (m *menu) NotFound(rw http.ResponseWriter, req *http.Request) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
