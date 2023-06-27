package chef

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
	allChefsRegex = regexp.MustCompile(`^\/chef[\/]?$`)         // /chef or /chef/
	chefNameRegex = regexp.MustCompile(`^\/chef\/([A-Za-z]+)$`) // /chef/job, /chef/ebukaOdi
)

// Chef json Object
type chefJson struct {
	Name  string `json:"name"`
	About string `json:"about"`
	Image string `json:"image"`
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
		c.deleteChef(rw, req)

	case req.Method == http.MethodDelete && chefNameRegex.MatchString(req.URL.Path):
		c.deleteChefByName(rw, req)

	default:
		c.notImplemented(rw) // returns 501 Not Implemented
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
		rw.WriteHeader(http.StatusNotFound)                    // 404
		rw.Write([]byte(http.StatusText(http.StatusNotFound))) // NotFound

		return // exit function call

	}

	// encode to json and rw sends the json
	err := json.NewEncoder(rw).Encode(chefNames)

	// error handling
	if err != nil {
		log.Fatal("error encoding into json")
	}

}

// client deletes all chef data
func (c *chef) deleteChef(rw http.ResponseWriter, req *http.Request) {

	// delete all element by re-initializing to nil
	*c = nil

	utils.Delete(rw, req, &c)

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

			fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")

			return // exit function call
		}

	}

	rw.WriteHeader(http.StatusNotFound)                    // 404
	rw.Write([]byte(http.StatusText(http.StatusNotFound))) // NotFound

}

// sends message to client if resource does not exist or not implemented
func (c *chef) notImplemented(rw http.ResponseWriter, req *http.Request) {

	utils.NotImplemented(rw) // returns 501 Not Implemented

}
