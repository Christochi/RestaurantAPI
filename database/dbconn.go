package database

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func conn() *sql.DB {

	// construct db url
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", "password"),
		Host:   "localhost:5432",
		Path:   "postgres", // dbname (POSTGRES_DB)
	}

	urlValues := dsn.Query()            //Query parses RawQuery and returns the corresponding values
	urlValues.Add("sslmode", "disable") // disbale sslmode and add to the map of Values

	// database url : postgres://postgres:password@localhost:5432/postgres?sslmode=disable
	dsn.RawQuery = urlValues.Encode()

	// create db connection
	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("sql open", err)
	}

	return db
}
