package utils

import (
	"database/sql"
	"log"
	"os"
)

// open and execute SQL script
func CreateTables(filename string, db *sql.DB) {

	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(contents))
	if err != nil {
		log.Fatal(err)
	}

}
