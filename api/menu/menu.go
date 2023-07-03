package menu

import (
	"encoding/json"
	"net/http"
	"regexp"
	"restaurantapi/utils"
	"strings"
)

// pathnames for subroot in url endpoint
var (
	allMenuRegex  = regexp.MustCompile(`^\/menu[\/]?$`)                                          // /menu or /menu/
	mealTypeRegex = regexp.MustCompile(`^\/menu\/([A-Za-z]+)$`)                                  // /menu/<anymealtype> = /menu/dinner
	mealRegex     = regexp.MustCompile(`^\/menu\/(breakfast|lunch|drinks|dinner)\/([A-Za-z]+)$`) // /menu/<anymealtype>/burger
)

// Menu json Object
type menuJson struct {
	Type  string `json:"type"` // meal type
	Meal  string `json:"meal"`
	Price string `json:"price"`
	Desc  string `json:"desc"` // description
}

type menu []menuJson // slice type to be used as a receiver for methods

// retuns a pointer to menu object
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
		m.postMenu(rw, req)

	case req.Method == http.MethodGet && allMenuRegex.MatchString(req.URL.Path):
		m.getMenu(rw)

	case req.Method == http.MethodGet && mealTypeRegex.MatchString(req.URL.Path):
		m.getMealType(rw, req)

	case req.Method == http.MethodGet && mealRegex.MatchString(req.URL.Path):
		m.getMeal(rw, req)

	case req.Method == http.MethodDelete && allMenuRegex.MatchString(req.URL.Path):
		m.deleteMenu(rw)

	case req.Method == http.MethodDelete && mealRegex.MatchString(req.URL.Path):
		m.deleteMeal(rw, req)

	default:
		http.Error(rw, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented) // returns 501 Not Implemented
	}

}

// client send menu data using POST Method
func (m *menu) postMenu(rw http.ResponseWriter, req *http.Request) {

	// read and decode to struct
	utils.Post(rw, req, m)

}

// client requests for menu data using GET Method
func (m *menu) getMenu(rw http.ResponseWriter) {

	// read and encode to json
	utils.Get(rw, m)

}

// gets the list of menu for a meal type
func (m *menu) getMealType(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := mealTypeRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /menu/<mealtype> = ["/menu/breakfast", "breakfast"]
	mealTypeName := strings.ToLower(urlSubPaths[1])

	var meal []menuJson // new slice to hold the filtered data

	for _, value := range *m {

		if strings.ToLower(value.Type) == mealTypeName {
			meal = append(meal, value) // append to new slice
		}

	}

	if meal == nil {
		rw.WriteHeader(http.StatusNotFound)                    // 404
		rw.Write([]byte(http.StatusText(http.StatusNotFound))) // NotFound

		return // exit function call
	}

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(meal)

	// error handling
	if err != nil {
		// 422 Unprocessable Entity
		http.Error(rw, "error encoding into json", http.StatusUnprocessableEntity)
	}

}

// gets a specific meal
func (m *menu) getMeal(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := mealRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /menu/<mealtype>/<mealname> = ["/menu/drinks/mangolasse", "drinks", "mangolasse"]
	mealName := strings.ToLower(urlSubPaths[2])

	var meal []menuJson // new slice to hold the filtered data

	for _, value := range *m {

		// remove whitespaces and returns lower case of the string
		name := strings.ToLower(strings.ReplaceAll(value.Meal, " ", ""))

		// compares if 2 strings have the same string literal or a substring
		if name == mealName || strings.Contains(name, mealName) {
			meal = append(meal, value) // append to new slice
		}

	}

	if meal == nil {
		rw.WriteHeader(http.StatusNotFound)                    // 404
		rw.Write([]byte(http.StatusText(http.StatusNotFound))) // NotFound

		return // exit function call
	}

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(meal)

	// error handling
	if err != nil {
		http.Error(rw, "error encoding into json", http.StatusUnprocessableEntity) // 422 Unprocessable Entity
	}

}

// client deletes all menu data
func (m *menu) deleteMenu(rw http.ResponseWriter) {

	// delete all element by re-initializing to nil
	*m = nil

	utils.Delete(rw, m)

}

// deletes a specific meal
func (m *menu) deleteMeal(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := mealRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the third index
	// example: /menu/<mealtype>/<mealname> = ["/menu/lunch/burger", "lunch", "burger"]
	meal := strings.ToLower(urlSubPaths[2])

	for index, value := range *m {

		// remove whitespaces and returns lower case of the string
		if strings.ToLower(strings.ReplaceAll(value.Meal, " ", "")) == meal {
			// delete an element
			(*m)[index] = (*m)[len(*m)-1] // replace the element with the last element
			*m = (*m)[:len(*m)-1]         // reinitialize the array with all the elements excluding last element

			rw.WriteHeader(http.StatusOK) // 200 OK
			rw.Write([]byte("resource deleted successfully"))

			return // exit function call
		}

	}

	rw.WriteHeader(http.StatusNotFound)                    // 404
	rw.Write([]byte(http.StatusText(http.StatusNotFound))) // NotFound

}
