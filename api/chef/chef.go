package chef

import (
	"encoding/json"
	"net/http"
	"regexp"
	"restaurantapi/utils"
	"strings"
)

// pathnames for subroot in url endpoint
var (
	allChefsRegex = regexp.MustCompile(`^\/chef[\/]?$`)         // /chef or /chef/
	chefNameRegex = regexp.MustCompile(`^\/chef\/([A-Za-z]+)$`) // /chef/job, /chef/ebukaOdi
)

// Chef json Object
type chefJson struct {
	Name  string `json:"name"`
	About string `json:"about"`
}

type chef []chefJson // slice type to be used as a receiver for methods

// retuns a pointer to chef object
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
		c.getChefs(rw)

	case req.Method == http.MethodGet && chefNameRegex.MatchString(req.URL.Path):
		c.getChefByName(rw, req)

	case req.Method == http.MethodPost && allChefsRegex.MatchString(req.URL.Path):
		c.postChef(rw, req)

	case req.Method == http.MethodDelete && allChefsRegex.MatchString(req.URL.Path):
		c.deleteChef(rw)

	case req.Method == http.MethodDelete && chefNameRegex.MatchString(req.URL.Path):
		c.deleteChefByName(rw, req)

	default:
		utils.ServerMessgae(rw, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented) // returns 501 Not Implemented
	}

}

// client send chef data using POST Method
func (c *chef) postChef(rw http.ResponseWriter, req *http.Request) {

	// read and decode to struct
	utils.Post(rw, req, c)

}

// client requests for chef data using GET Method
func (c *chef) getChefs(rw http.ResponseWriter) {

	// read and encode to json
	utils.Get(rw, c)

}

// client requests for specific chef
func (c *chef) getChefByName(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := chefNameRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /chef/<name> = ["/chef/stevejob", "stevejobs"]
	name := strings.ToLower(urlSubPaths[1])

	var chefNames []chefJson // new slice to hold the filtered data

	for _, value := range *c {

		// remove whitespaces and returns lower case of the string
		subPath := strings.ToLower(strings.ReplaceAll(value.Name, " ", ""))

		// compares if 2 strings have the same string literal or a substring
		if subPath == name || strings.Contains(subPath, name) {
			chefNames = append(chefNames, value) // append to new slice
		}

	}

	if chefNames == nil {
		utils.ServerMessgae(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found

		return // exit function call

	}

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(chefNames)

	// error handling
	if err != nil {
		http.Error(rw, "error encoding into json", http.StatusUnprocessableEntity) // 422 Unprocessable Entity
	}

}

// client deletes all chef data
func (c *chef) deleteChef(rw http.ResponseWriter) {

	// delete all element by re-initializing to nil
	*c = nil

	utils.Delete(rw, c)

}

// client deletes a specific chef
func (c *chef) deleteChefByName(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := chefNameRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /user/SaintLawrence = ["/user/SaintLawrence", "SaintLawrence"]
	name := strings.ToLower(urlSubPaths[1])

	for index, value := range *c {

		// remove whitespaces and returns lower case of the string
		if strings.ToLower(strings.ReplaceAll(value.Name, " ", "")) == name {
			// delete an element
			(*c)[index] = (*c)[len(*c)-1] // replace the element with the last element
			*c = (*c)[:len(*c)-1]         // reinitialize the array with all the elements excluding last element

			utils.ServerMessgae(rw, "resource deleted successfully", http.StatusOK) // 200 OK

			return // exit function call
		}

	}

	utils.ServerMessgae(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found
}
