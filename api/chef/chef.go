package chef

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	errs "restaurantapi/errors"
	"restaurantapi/utils"
	"strings"

	"github.com/Christochi/error-handler/service"
)

var logger = utils.InfoLog() // return info field

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

	case req.Method == http.MethodPut && chefNameRegex.MatchString(req.URL.Path):
		c.putChef(rw, req)

	case req.Method == http.MethodDelete && allChefsRegex.MatchString(req.URL.Path):
		c.deleteChef(rw)

	case req.Method == http.MethodDelete && chefNameRegex.MatchString(req.URL.Path):
		c.deleteChefByName(rw, req)

	default:
		utils.ServerMessage(rw, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented) // returns 501 Not Implemented
	}

}

// Insert into the chef table unique values (no duplicates)
func (c *chef) bulkInsert(query string, db *sql.DB) (int64, error) {

	var row, numOfRows int64 //  db row affected and the total number of rows affected
	var result sql.Result    // interface for invoking RowsAffected()
	var err error

	for _, column := range *c {

		// insert into table
		result, err = utils.Database.Exec(query, column.Name, column.About, column.Image, column.Gender, column.Age)
		if err != nil {
			return 0, service.NewError(err, "Can't insert into table")
		}

		// return number of table rows with inserted data
		row, err = result.RowsAffected()
		if err != nil {
			return 0, service.NewError(err, "no table rows were affected")
		}

		numOfRows += row // increment

	}

	return numOfRows, nil

}

// traverse the table rows
func (c *chef) iterDBRows(rows *sql.Rows, column chefJson) error {

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&column.Name, &column.About, &column.Image, &column.Gender, &column.Age); err != nil {
			return service.NewError(err, "could not fetch table rows")
		}

		*c = append(*c, column)
	}

	if err := rows.Err(); err != nil {
		return service.NewError(err, "error occured during iteration")
	}

	return nil

}

// client send chef data using POST Method
func (c *chef) postChef(rw http.ResponseWriter, req *http.Request) {

	// log for informational purpose
	logger.Println("POST chef request at /chef endpoint")

	// read and decode to struct
	err := utils.Create(rw, req, c)
	if err != nil {
		logger.Println(errs.RestError(rw, err))
		return // exit method to avoid superflous call to response.WriteHeader
	}

	if utils.Database != nil {

		// Delete all rows from the chef table since it is a POST request
		// and reset PK to 1
		err := utils.ExecuteQueries(utils.DeleteChefRowsQuery, utils.Database)
		if err != nil {
			logger.Println(errs.DatabaseError(err))
		}

		// Insert into the chef table unique values (no duplicates) and store the number of table rows affected
		numOfRows, err := c.bulkInsert(utils.ChefBulkInsertQuery, utils.Database)
		if err != nil {
			logger.Println(errs.DatabaseError(err))
		}

		message := fmt.Sprintf("%d row(s) in the table were created", numOfRows) // construct server message
		utils.ServerMessage(rw, message, http.StatusCreated)                     // send server response

	}

}

// client requests for chef data using GET Method
func (c *chef) getChefs(rw http.ResponseWriter) {

	// log for informational purpose
	logger.Println("GET chef request at /chef endpoint")

	if utils.Database != nil {

		// initialize to nil to clear any initial value so that fresh copy of the data in db can be stored
		*c = nil

		var column chefJson // placeholder for column values

		// get the rows from table
		rows, err := utils.SelectRows(utils.SelectAllChefsQuery, utils.Database)
		if err != nil {
			logger.Println(errs.DatabaseError(err))
		}

		// iterate table
		err = c.iterDBRows(rows, column)
		if err != nil {
			logger.Println(errs.DatabaseError(err))
		}

	}

	// read and encode to json
	err := utils.Get(rw, c)
	if err != nil {
		logger.Println(errs.RestError(rw, err))
	}

}

// client requests for specific chef
func (c *chef) getChefByName(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := chefNameRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /chef/<name> = ["/chef/stevejobs", "stevejobs"]
	name := strings.ToLower(urlSubPaths[1])

	// log for informational purpose
	logger.Printf("GET chef name request at /chef/%s endpoint", name)

	if utils.Database != nil {

		// initialize to nil to clear any initial value so that fresh copy of the data in db can be stored
		*c = nil

		var column chefJson // placeholder for column values

		// Retrieve data from the table that matches the substring
		rows, err := utils.SelectRows(utils.SelectChefByNameQuery, utils.Database, name+"%")
		if err != nil {
			logger.Println(errs.DatabaseError(err))
		}

		// iterate table
		err = c.iterDBRows(rows, column)
		if err != nil {
			logger.Println(errs.DatabaseError(err))
		}

		if *c == nil {
			utils.ServerMessage(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found
			return                                                                             // exit function call
		}

	}

	// read and encode to json
	err := utils.Get(rw, c)
	if err != nil {
		logger.Println(errs.RestError(rw, err))
	}

}

// Update or Create a chef and store in table
func (c *chef) putChef(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := chefNameRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /chef/<name> = ["/chef/stevejobs", "stevejobs"]
	name := strings.ToLower(urlSubPaths[1])

	// log for informational purpose
	logger.Printf("PUT chef request at /chef/%s endpoint", name)

	// if utils.Database != nil {

	// 	*c = nil

	// 	var column chefJson // placeholder for column values

	// 	// fetch the value of image_name column since it is unique
	// 	query := `SELECT full_name, about, image_name, gender, age from chef where LOWER(REPLACE(full_name, ' ', '')) = $1;`
	// 	rows := utils.SelectRows(query, utils.Database, name)
	// 	c.iterDBRows(rows, column)

	// 	// store image name
	// 	var imageName string
	// 	for _, col := range *c {
	// 		imageName = col.Image
	// 	}

	// 	// read and decode to struct
	// 	apiErr := utils.Create(rw, req, c)
	// 	if apiErr != nil {
	// 		errs.RestError(rw, apiErr)
	// 		return // exit method to avoid superflous call to response.WriteHeader
	// 	}

	// 	// Update row if it exist, else Create new row
	// 	var numOfRows int64   //  db row affected
	// 	var result sql.Result // interface for invoking RowsAffected()
	// 	var err error

	// 	for _, column := range *c {

	// 		// insert into table
	// 		result, err = utils.Database.Exec(utils.UpdateAChef, imageName, column.Name, column.About, column.Image, column.Gender, column.Age)
	// 		if err != nil {
	// 			log.Fatal("Exec, ", err)
	// 		}

	// 		// return number of table rows with inserted data
	// 		numOfRows, err = result.RowsAffected()
	// 		if err != nil {
	// 			log.Fatal("Result err, ", err)
	// 		}

	// 	}
	// 	// Insert into the chef table unique values (no duplicates) and store the number of table rows affected
	// 	// numOfRows := c.bulkInsert(utils.ChefBulkInsertQuery, utils.Database)

	// 	message := fmt.Sprintf("%d row(s) in the table have been updated or created", numOfRows) // construct server message
	// 	utils.ServerMessage(rw, message, http.StatusCreated)                                     // send server response}

	// }

}

// client deletes all chef data
func (c *chef) deleteChef(rw http.ResponseWriter) {

	// log for informational purpose
	logger.Println("DELETE chef request at /chef endpoint")

	// delete all element by re-initializing to nil
	*c = nil

	if utils.Database != nil {
		// Delete all rows from the chef table and reset PK to 1
		err := utils.ExecuteQueries(utils.DeleteChefRowsQuery, utils.Database)
		if err != nil {
			logger.Println(errs.DatabaseError(err))
		}

		utils.ServerMessage(rw, "table row(s) deleted successfully", http.StatusOK) // 200 OK
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
	logger.Printf("DELETE chef name request at /chef/%s endpoint", name)

	if utils.Database != nil {
		// Delete a row from the chef table
		result, err := utils.Database.Exec(utils.DeleteAChefQuery, name)
		if err != nil {
			err = service.NewError(err, "error in delete statement")
			logger.Println(errs.DatabaseError(err))
		}

		// return number of rows deleted
		numOfRows, err := result.RowsAffected()
		if err != nil {
			err = service.NewError(err, "no rows found")
			logger.Println(errs.DatabaseError(err))
		}

		if numOfRows > 0 {
			utils.ServerMessage(rw, "table row(s) deleted successfully", http.StatusOK) // 200 OK
		} else {
			utils.ServerMessage(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found
		}
	}

}
