package database

import (
	"fmt"
	"testing"
)

// Test DB Connection
func TestConn(t *testing.T) {

	connTest := conn()     // returns the database
	err := connTest.Ping() // ping db
	if err != nil {
		fmt.Println("test conn", err)
	}
}
