package chef

import (
	"database/sql"
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

// Insert into the chef table unique values (no duplicates)
func (c *chef) bulkInsert(query string, db *sql.DB) {

	for _, column := range *c {

		_, err := utils.Database.Exec(query, column.Name, column.About, column.Image, column.Gender, column.Age)
		if err != nil {
			log.Fatal("Exec, ", err)
		}

	}

}

// traverse the db rows
func (c *chef) iterDBRows(rows *sql.Rows, column chefJson) {

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&column.Name, &column.About, &column.Image, &column.Gender, &column.Age); err != nil {
			log.Fatal("Scan error, ", err)
		}

		*c = append(*c, column)

	}

}

// client send chef data using POST Method
func (c *chef) postChef(rw http.ResponseWriter, req *http.Request) {

	// log for informational purpose
	requestLogger.Println("POST chef request at /chef endpoint")

	// read and decode to struct
	utils.Post(rw, req, c)

	// Delete all rows from the chef table since it is a POST request
	// and reset PK to 1
	utils.ExecuteQueries(utils.DeleteChefRowsQuery, utils.Database)

	// Insert into the chef table unique values (no duplicates)
	c.bulkInsert(utils.ChefBulkInsertQuery, utils.Database)

}

// client requests for chef data using GET Method
func (c *chef) getChefs(rw http.ResponseWriter) {

	// initialize to nil to clear any initial value so that fresh copy of the data in db can be stored
	*c = nil

	// log for informational purpose
	requestLogger.Println("GET chef request at /chef endpoint")

	var column chefJson // placeholder for column values

	// get the rows from db
	rows := utils.SelectRows(utils.SelectAllChefsQuery, utils.Database)
	c.iterDBRows(rows, column)

	// read and encode to json
	utils.Get(rw, c)

}

// client requests for specific chef
func (c *chef) getChefByName(rw http.ResponseWriter, req *http.Request) {

	// initialize to nil to clear any initial value so that fresh copy of the data in db can be stored
	*c = nil

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := chefNameRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /chef/<name> = ["/chef/stevejobs", "stevejobs"]
	name := strings.ToLower(urlSubPaths[1])

	// log for informational purpose
	requestLogger.Printf("GET chef name request at /chef/%s endpoint", name)

	var column chefJson // placeholder for column values

	// Retrieve data that matches the substring from the db
	rows := utils.SelectRows(utils.SelectChefByNameQuery, utils.Database, name+"%")
	c.iterDBRows(rows, column)

	if *c == nil {
		utils.ServerMessage(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found

		return // exit function call

	}

	// read and encode to json
	utils.Get(rw, c)

}

// client deletes all chef data
func (c *chef) deleteChef(rw http.ResponseWriter) {

	// log for informational purpose
	requestLogger.Println("DELETE chef request at /chef endpoint")

	// delete all element by re-initializing to nil
	*c = nil

	// Delete all rows from the chef table and reset PK to 1
	utils.ExecuteQueries(utils.DeleteChefRowsQuery, utils.Database)

	utils.ServerMessage(rw, "resource deleted successfully", http.StatusOK) // 200 OK

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

	// Delete a row from the chef table
	result, err := utils.Database.Exec(utils.DeleteAChefQuery, name)
	if err != nil {
		log.Fatal("Exec err, ", err)
	}

	numOfRoles, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Result err, ", err)
	}

	if numOfRoles > 0 {
		utils.ServerMessage(rw, "resource deleted successfully", http.StatusOK) // 200 OK
	} else {
		utils.ServerMessage(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found
	}

	// for index, value := range *c {

	// 	// remove whitespaces and returns lower case of the string
	// 	if strings.ToLower(strings.ReplaceAll(value.Name, " ", "")) == name {
	// 		// delete an element
	// 		(*c)[index] = (*c)[len(*c)-1] // replace the element with the last element
	// 		*c = (*c)[:len(*c)-1]         // reinitialize the array with all the elements excluding last element

	// 		utils.ServerMessage(rw, "resource deleted successfully\n", http.StatusOK) // 200 OK

	// 		// Delete all rows from the chef table and reset PK to 1
	// 		utils.ExecuteQueries(utils.DeleteChefRowsQuery, utils.Database)

	// 		c.bulkInsert(utils.ChefBulkInsertQuery, utils.Database) // bulk insert into db

	// 		fmt.Fprint(rw, "created db rows") // 200 OK

	// 		return // exit function call
	// 	}

	// }

}
