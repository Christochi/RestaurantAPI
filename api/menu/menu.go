package menu

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"restaurantapi/utils"
	"strings"
)

var requestLogger = utils.InfoLog() // return info field

// pathnames for subroot in url endpoint
var (
	allMenuRegex  = regexp.MustCompile(`^\/menu[\/]?$`)                                          // /menu or /menu/
	mealTypeRegex = regexp.MustCompile(`^\/menu\/(breakfast|lunch|drinks|dinner)$`)              // /menu/<anymealtype> = /menu/dinner
	mealRegex     = regexp.MustCompile(`^\/menu\/(breakfast|lunch|drinks|dinner)\/([A-Za-z]+)$`) // /menu/<anymealtype>/burger
)

// Menu json Object
type menuJson struct {
	Type      string `json:"type"` // meal type
	Meal      string `json:"meal"`
	Price     string `json:"price"`
	Desc      string `json:"desc"`  // description
	Image     string `json:"image"` // image name
	Available bool   `json:"available"`
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
		utils.ServerMessage(rw, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented) // returns 501 Not Implemented
	}

}

// Insert into the menu table unique values (no duplicates)
func (m *menu) bulkInsert(query string, db *sql.DB) int64 {

	var row, numOfRows int64 //  db row affected and the total number of rows affected
	var result sql.Result    // interface for invoking RowsAffected()
	var err error

	// insert into table
	for _, column := range *m {

		result, err = utils.Database.Exec(query, column.Type, column.Meal, column.Price, column.Desc, column.Image, column.Available)
		if err != nil {
			log.Fatal("Exec, ", err)
		}

		// return number of table rows with inserted data
		row, err = result.RowsAffected()
		if err != nil {
			log.Fatal("Result err, ", err)
		}

		numOfRows += row // increment

	}

	return numOfRows

}

// traverse the db rows
func (m *menu) iterDBRows(rows *sql.Rows, column menuJson) {

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&column.Type, &column.Meal, &column.Price, &column.Desc, &column.Image, &column.Available); err != nil {
			log.Fatal("Scan error, ", err)
		}

		*m = append(*m, column)

	}

}

// client send menu data using POST Method
func (m *menu) postMenu(rw http.ResponseWriter, req *http.Request) {

	// log for informational purpose
	requestLogger.Println("POST menu request at /menu endpoint")

	// read and decode to struct
	utils.Create(rw, req, m)

	if utils.Database != nil {
		// Delete all rows from the menu table since it is a POST request
		// and reset PK to 1
		utils.ExecuteQueries(utils.DeleteMenuRowsQuery, utils.Database)

		// Insert into the menu table unique values (no duplicates) and store the number of table rows affected
		numOfRows := m.bulkInsert(utils.MenuBulkInsertQuery, utils.Database)

		message := fmt.Sprintf("%d row(s) in the table were created", numOfRows) // construct server message
		utils.ServerMessage(rw, message, http.StatusCreated)                     // send server response
	}

}

// client requests for menu data using GET Method
func (m *menu) getMenu(rw http.ResponseWriter) {

	// log for informational purpose
	requestLogger.Println("GET menu request at /menu endpoint")

	if utils.Database != nil {

		// initialize to nil to clear any initial value so that fresh copy of the data in db can be stored
		*m = nil

		var column menuJson // placeholder for column values

		// get the rows from table
		rows := utils.SelectRows(utils.SelectAllMenuQuery, utils.Database)
		m.iterDBRows(rows, column)

	}

	// read and encode to json
	utils.Get(rw, m)

}

// gets the list of menu for a meal type
func (m *menu) getMealType(rw http.ResponseWriter, req *http.Request) {

	// initialize to nil to clear any initial value so that fresh copy of the data in db can be stored
	*m = nil

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := mealTypeRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /menu/<mealtype> = ["/menu/breakfast", "breakfast"]
	mealTypeName := strings.ToLower(urlSubPaths[1])

	// log for informational purpose
	requestLogger.Printf("GET meal type request at /menu/%s endpoint", mealTypeName)

	var column menuJson // placeholder for column values

	// Retrieve data from the table that matches the substring and append to menu struct
	rows := utils.SelectRows(utils.SelectMealTypeQuery, utils.Database, mealTypeName)
	m.iterDBRows(rows, column)

	if *m == nil {
		utils.ServerMessage(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found

		return // exit function call
	}

	// read and encode to json
	utils.Get(rw, m)

}

// gets a specific meal
func (m *menu) getMeal(rw http.ResponseWriter, req *http.Request) {

	// initialize to nil to clear any initial value so that fresh copy of the data in db can be stored
	*m = nil

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := mealRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the second index
	// example: /menu/<mealtype>/<mealname> = ["/menu/drinks/mangolasse", "drinks", "mangolasse"]
	mealTypeName := strings.ToLower(urlSubPaths[1])
	mealName := strings.ToLower(urlSubPaths[2])

	// log for informational purpose
	requestLogger.Printf("GET meal request at /menu/%s/%s endpoint", mealTypeName, mealName)

	// var meal []menuJson // new slice to hold the filtered data
	var column menuJson // placeholder for column values

	// Retrieve data from the table that matches the substring
	rows, err := utils.Database.Query(utils.SelectMealQuery, mealTypeName, "%"+mealName+"%")
	if err != nil {
		log.Fatal("Select error, ", err)
	}

	m.iterDBRows(rows, column) // traverse the db rows and append to menu struct

	if *m == nil {
		utils.ServerMessage(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found

		return // exit function call
	}

	// read and encode to json
	utils.Get(rw, m)

}

// client deletes all menu data
func (m *menu) deleteMenu(rw http.ResponseWriter) {

	// log for informational purpose
	requestLogger.Println("DELETE menu request at /menu endpoint")

	// delete all element by re-initializing to nil
	*m = nil

	if utils.Database != nil {
		// Delete all rows from the menu table and reset PK to 1
		utils.ExecuteQueries(utils.DeleteMenuRowsQuery, utils.Database)

		utils.ServerMessage(rw, "table row(s) deleted successfully", http.StatusOK) // 200 OK
	}

}

// deletes a specific meal
func (m *menu) deleteMeal(rw http.ResponseWriter, req *http.Request) {

	// returns slice of substrings that matches subexpressions in the url
	urlSubPaths := mealRegex.FindStringSubmatch(req.URL.Path)

	// since the order of the slice is known, store the third index
	// example: /menu/<mealtype>/<mealname> = ["/menu/lunch/burger", "lunch", "burger"]
	mealTypeName := strings.ToLower(urlSubPaths[1])
	meal := strings.ToLower(urlSubPaths[2])

	// log for informational purpose
	requestLogger.Printf("DELETE meal request at /menu/%s/%s endpoint", mealTypeName, meal)

	if utils.Database != nil {
		// Delete a row from the menu table
		result, err := utils.Database.Exec(utils.DeleteAMealQuery, mealTypeName, "%"+meal+"%")
		if err != nil {
			log.Fatal("Exec err, ", err)
		}

		// return number of rows deleted
		numOfRoles, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Result err, ", err)
		}

		if numOfRoles > 0 {
			utils.ServerMessage(rw, "table row deleted successfully", http.StatusOK) // 200 OK
		} else {
			utils.ServerMessage(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404 Not Found
		}

	}

}
