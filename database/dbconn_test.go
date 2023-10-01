package database

import (
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

// Test DB Connection
func TestConn(t *testing.T) {

	// load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	connTest := Conn()     // returns the database
	err := connTest.Ping() // ping db
	if err != nil {
		fmt.Println("test conn", err)
	}

	defer connTest.Close() // close db
}
