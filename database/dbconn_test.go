package database

import (
	"log"
	errs "restaurantapi/errors"
	"testing"

	"github.com/Christochi/error-handler/service"
	"github.com/joho/godotenv"
)

// Test DB Connection
func TestConn(t *testing.T) {

	// load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	connTest, err := Conn() // returns the database

	// validate if a db driver was supplied
	if err != nil {
		log.Fatal(errs.DatabaseError(err))
	}

	err = connTest.Ping() // ping db

	// validate DSN
	if err != nil {
		err = service.NewError(err, "invalid data source name")
		log.Fatal(errs.DatabaseError(err))
	}

	defer connTest.Close() // close db
}
