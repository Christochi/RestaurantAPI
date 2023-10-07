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

// BULK SELECTION
func BulkSelect(query string, db *sql.DB) *sql.Rows {

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Select error, ", err)
	}

	return rows

}
