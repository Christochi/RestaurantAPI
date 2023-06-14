package menu

import (
	"regexp"
)

// pathnames for subroot in url endpoint
var (
	allMenuRegex = regexp.MustCompile(`^\/menu[\/]?$`)
)
