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
	allChefsRegex     = regexp.MustCompile(`^\/chef[\/]?$`)         // /chef or /chef/
	specificChefRegex = regexp.MustCompile(`^\/chef\/([A-Za-z]+)$`) // /chef/job, /chef/Cynthia
)

// Chef json Object
type chefJson struct {
	Name  string `json:"name"`
	About string `json:"about"`
}

type chef []chefJson // slice type to be used as a receiver for methods

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
		c.getChef(rw, req)

	case req.Method == http.MethodGet && specificChefRegex.MatchString(req.URL.Path):
		c.GetChefByName(rw, req)

	case req.Method == http.MethodPost && allChefsRegex.MatchString(req.URL.Path):
		c.postChef(rw, req)

	case req.Method == http.MethodDelete && allChefsRegex.MatchString(req.URL.Path):
		c.DeleteChef(rw, req)

	case req.Method == http.MethodDelete && specificChefRegex.MatchString(req.URL.Path):
		c.DeleteChefByName(rw, req)

	default:
		c.notFound(rw, req) // returns 501 Not Implemented
	}

}

// client send chef data using POST Method
func (c *chef) postChef(rw http.ResponseWriter, req *http.Request) {

	// read and decode to struct
	utils.Post[*chef](rw, req, c)

}

// client requests for chef data using GET Method
func (c *chef) getChef(rw http.ResponseWriter, req *http.Request) {

	// read and encode to json
	utils.Get[*chef](rw, req, c)

}

// client requests for specific chef
func (c *chef) GetChefByName(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := specificChefRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /chef/<name> = ["/chef/job", "job"]
	name := urlSubPaths[1]

	var chefNames []chefJson // new slice to hold the filtered data

	for _, value := range *c {

		// remove whitespaces and returns lower case of the string
		if strings.ToLower(strings.ReplaceAll(value.Name, " ", "")) == name {
			chefNames = append(chefNames, value) // append to new slice

			// encode to json and rw sends the json
			err := json.NewEncoder(rw).Encode(chefNames)

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

// client deletes all chef data
func (c *chef) DeleteChef(rw http.ResponseWriter, req *http.Request) {

	// delete all element by re-initializing to nil
	*c = nil

	if *c != nil {
		fmt.Fprintf(rw, "%s\n", http.StatusText(http.StatusInternalServerError)) // 500 Internal Server Error
	} else {
		fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")
	}

}

// client deletes a specific chef
func (c *chef) DeleteChefByName(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := specificChefRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /user/job = ["/user/job", "job"]
	name := urlSubPaths[1]

	for index, value := range *c {

		// remove whitespaces and returns lower case of the string
		if strings.ToLower(strings.ReplaceAll(value.Name, " ", "")) == name {
			// delete an element
			(*c)[index] = (*c)[len(*c)-1] // replace the element with the last element
			*c = (*c)[:len(*c)-1]         // reinitialize the array with all the elements excluding last element

			fmt.Fprintln(rw, http.StatusOK, http.StatusText(http.StatusOK), "resource deleted successfully")

			return
		}

	}

	rw.WriteHeader(http.StatusNotFound)                    // 404
	rw.Write([]byte(http.StatusText(http.StatusNotFound))) // NotFound

}

// sends message to client if resource does not exist or not implemented
func (c *chef) notFound(rw http.ResponseWriter, req *http.Request) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
