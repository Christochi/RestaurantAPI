package menu

import (
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
