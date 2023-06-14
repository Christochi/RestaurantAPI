package menu

import (
	"net/http"
	"regexp"
)

// pathnames for subroot in url endpoint
var (
	allMenuRegex = regexp.MustCompile(`^\/menu[\/]?$`)
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
func NewChef() *menu {

	return new(menu)

}

// handlerfunc for menu endpoint
func (c *menu) MenuHandler(rw http.ResponseWriter, req *http.Request) {

	// inform browser to expect json
	rw.Header().Set("Content-Type", "application/json")

	// determines the function to call by the request
	switch {

	case req.Method == http.MethodPost && allMenuRegex.MatchString(req.URL.Path):
		c.PostMenu(rw, req)

	// case req.Method == http.MethodGet && allChefsRegex.MatchString(req.URL.Path):
	// 	c.GetChef(rw, req)

	// case req.Method == http.MethodGet && specificChefRegex.MatchString(req.URL.Path):
	// 	c.GetChefByName(rw, req)

	// case req.Method == http.MethodDelete && allChefsRegex.MatchString(req.URL.Path):
	// 	c.DeleteChef(rw, req)

	// case req.Method == http.MethodDelete && specificChefRegex.MatchString(req.URL.Path):
	// 	c.DeleteChefByName(rw, req)

	default:
		c.notFound(rw, req) // returns 501 Not Implemented
	}

}

// sends message to client if request does not exist or not implemented
func (c *menu) notFound(rw http.ResponseWriter, req *http.Request) {

	rw.WriteHeader(http.StatusNotImplemented)                    // 501
	rw.Write([]byte(http.StatusText(http.StatusNotImplemented))) // Not Implemented

}
