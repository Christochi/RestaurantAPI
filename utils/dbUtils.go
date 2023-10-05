package utils

import (
	"database/sql"
	"log"
	"os"
)

var Database *sql.DB // place holder for the database

const (
	ChefRowsDeleteQuery = `DELETE FROM chef; 
   	ALTER SEQUENCE chef_id_seq RESTART WITH 1;`

	ChefBulkInsertQuery = `INSERT INTO chef (full_name, about, image_name, gender, age) 
		VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING;`
)

// open and read SQL script
func ReadSQLScript(filename string) []byte {

	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("ReadFile error, ", err)
	}

	return contents

}

// execute SQL queries
func ExecuteQueries(query string, db *sql.DB) {

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Exec Queries, ", err)
	}

}
