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

	ChefBulkInsertQuery = `INSERT INTO chef (full_name, about, image_name, gender, age) 
		VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING;`

	SelectAllChefRowsQuery = `SELECT full_name, about, image_name, gender, age FROM chef;`

	SelectAnyChefRowsQuery = `SELECT full_name, about, image_name, gender, age FROM chef WHERE full_name LIKE $1;`
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

// Return DB rows
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
	argsExist := func() *sql.Rows {

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

	if len(args) >= 1 {
		return argsExist()
	} else {
		return noArgs()
	}

}
