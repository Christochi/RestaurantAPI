package utils

import (
	"database/sql"
	"fmt"
	"io"
	"os"
)

// open and execute SQL script
func OpenFile(filename string, db *sql.DB) {

	file, err := os.Open(filename)
	if err == nil {
		panic(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	execSQLScript(content, db)

}

func execSQLScript(content []byte, db *sql.DB) {

	_, err := db.Exec(string(content))
	if err != nil {
		fmt.Println(err)
		return
	}

}
