package utils

import (
	"database/sql"
	"log"
	"os"
)

var Database *sql.DB // place holder for the database

// SQL Queries
const (
	DeleteChefRowsQuery = `DELETE FROM chef; 
   	ALTER SEQUENCE chef_id_seq RESTART WITH 1;`
	DeleteMenuRowsQuery = `DELETE FROM menu; 
		ALTER SEQUENCE menu_id_seq RESTART WITH 1;`

	DeleteAChefQuery = `DELETE FROM chef WHERE LOWER(REPLACE(full_name, ' ', '')) = $1`
	DeleteAMealQuery = `DELETE FROM MENU WHERE LOWER(meal_type) = $1 AND LOWER(REPLACE(meal_name, ' ', '')) LIKE $2;`

	ChefBulkInsertQuery = `INSERT INTO chef (full_name, about, image_name, gender, age) 
		VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING;`
	MenuBulkInsertQuery = `INSERT INTO menu (meal_type, meal_name, price, about, image_name, available) 
		VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING;`

	SelectAllChefsQuery = `SELECT full_name, about, image_name, gender, age FROM chef;`
	SelectAllMenuQuery  = `SELECT meal_type, meal_name, price, about, image_name, available FROM menu;`

	SelectChefByNameQuery = `SELECT full_name, about, image_name, gender, age FROM chef WHERE full_name LIKE INITCAP($1);`
	SelectMealTypeQuery   = `SELECT meal_type, meal_name, price, about, image_name, available FROM menu WHERE meal_type = INITCAP($1);`
	SelectMealQuery       = `SELECT meal_type, meal_name, price, about, image_name, available FROM menu WHERE LOWER(meal_type) = $1 
		AND LOWER(REPLACE(meal_name, ' ', '')) LIKE $2;`
)

// open and read SQL script
func ReadSQLScript(filename string) []byte {

	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("ReadFile error, ", err)
	}

	return contents

}

// execute SQL queries - CREATE, INSERT, UPDATE & DELETE
func ExecuteQueries(query string, db *sql.DB) {

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Exec Queries, ", err)
	}

}

// Return table rows
func SelectRows(query string, db *sql.DB, args ...string) *sql.Rows {

	// accepts no SQL placeholder argument
	noArgs := func() *sql.Rows {

		rows, err := db.Query(query)
		if err != nil {
			log.Fatal("Select error, ", err)
		}

		return rows

	}

	// accepts SQL placeholder arguments
	withArgs := func() *sql.Rows {

		var rows *sql.Rows

		for _, arg := range args {
			var err error

			rows, err = db.Query(query, arg)
			if err != nil {
				log.Fatal("Select2 error, ", err)
			}

		}
		return rows

	}

	if len(args) > 0 {
		return withArgs()
	} else {
		return noArgs()
	}

}
