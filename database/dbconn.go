package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Conn() *sql.DB {

	// returns hostname and port as a string
	host := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	// construct db url
	dsn := url.URL{
		Scheme: os.Getenv("URL_SCHEME"),
		User:   url.UserPassword(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD")),
		Host:   host,
		Path:   os.Getenv("POSTGRES_DB"),
	}

	urlValues := dsn.Query()                       //Query parses RawQuery and returns the corresponding values
	urlValues.Add("sslmode", os.Getenv("SSLMODE")) // disbale sslmode and add to the map of Values

	// database url = URL_SCHEME://POSTGRES_USER:POSTGRES_PASSWORD@HOST:PORT/POSTGRES_DB?sslmode=SSLMODE
	dsn.RawQuery = urlValues.Encode()

	// create db connection
	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		panic(err)
	}

	return db

}
