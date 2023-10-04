package chef

import (
	"log"
	"net/http"
	"regexp"
	"restaurantapi/utils"
	"strings"
)

var requestLogger = utils.InfoLog() // return info field

// pathnames for subroot in url endpoint
var (
	allChefsRegex = regexp.MustCompile(`^\/chef[\/]?$`)         // /chef or /chef/
	chefNameRegex = regexp.MustCompile(`^\/chef\/([A-Za-z]+)$`) // /chef/job, /chef/ebukaOdi
)

// Chef json Object
type chefJson struct {
	Name   string `json:"name"`
	About  string `json:"about"`
	Image  string `json:"image"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
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
		utils.ServerMessage(rw, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented) // returns 501 Not Implemented
	}

}

// client send chef data using POST Method
func (c *chef) postChef(rw http.ResponseWriter, req *http.Request) {

	// log for informational purpose
	requestLogger.Println("POST chef request at /chef endpoint")

	// read and decode to struct
	utils.Post(rw, req, c)

	query := `INSERT INTO chef (full_name, about, image_name, gender, age) 
					VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING;`

	// Insert into the chef table unique values (no duplicates)
	for _, elem := range *c {

		_, err := utils.Database.Exec(query, elem.Name, elem.About, elem.Image, elem.Gender, elem.Age)
		if err != nil {
			log.Fatal("Exec, ", err)
		}

	}

}

// client requests for chef data using GET Method
func (c *chef) getChefs(rw http.ResponseWriter) {

	// log for informational purpose
	requestLogger.Println("GET chef request at /chef endpoint")

	// read and encode to json
	utils.Get(rw, c)

}

// client requests for specific chef
func (c *chef) getChefByName(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := chefNameRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /chef/<name> = ["/chef/stevejobs", "stevejobs"]
	name := strings.ToLower(urlSubPaths[1])

	// log for informational purpose
	requestLogger.Printf("GET chef name request at /chef/%s endpoint", name)

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
		utils.ServerMessage(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found

		return // exit function call

	}

	// read and encode to json
	utils.Get(rw, chefNames)

}

// client deletes all chef data
func (c *chef) deleteChef(rw http.ResponseWriter) {

	// log for informational purpose
	requestLogger.Println("DELETE chef request at /chef endpoint")

	// delete all element by re-initializing to nil
	*c = nil

	utils.Delete(rw, c)

	// Delete all rows from the chef table and reset PK to 1
	query := `DELETE FROM chef;
	 			ALTER SEQUENCE chef_id_seq RESTART WITH 1;`

	_, err := utils.Database.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

// client deletes a specific chef
func (c *chef) deleteChefByName(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := chefNameRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /user/SaintLawrence = ["/user/SaintLawrence", "SaintLawrence"]
	name := strings.ToLower(urlSubPaths[1])

	// log for informational purpose
	requestLogger.Printf("DELETE chef name request at /chef/%s endpoint", name)

	for index, value := range *c {

		// remove whitespaces and returns lower case of the string
		if strings.ToLower(strings.ReplaceAll(value.Name, " ", "")) == name {
			// delete an element
			(*c)[index] = (*c)[len(*c)-1] // replace the element with the last element
			*c = (*c)[:len(*c)-1]         // reinitialize the array with all the elements excluding last element

			utils.ServerMessage(rw, "resource deleted successfully", http.StatusOK) // 200 OK

			// Delete all rows from the chef table and reset PK to 1
			// query := `DELETE FROM chef;
			// 			 ALTER SEQUENCE chef_id_seq RESTART WITH 1;`

			// _, err := utils.Database.Exec(query)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			// c.postChef(rw, req)

			return // exit function call
		}

	}

	utils.ServerMessage(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found
}
