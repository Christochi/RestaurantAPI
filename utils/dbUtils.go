package utils

import (
	"database/sql"
	"log"
	"os"
)

var Database *sql.DB // place holder for the database

// open and execute SQL script
func CreateTables(filename string, db *sql.DB) {

	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("ReadFile error, ", err)
	}

	_, err = db.Exec(string(contents))
	if err != nil {
		log.Fatal("Exec, ", err)
	}

}
