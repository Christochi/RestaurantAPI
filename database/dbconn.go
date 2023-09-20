package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func conn() *sql.DB {

	// load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// construct db url
	dsn := url.URL{
		Scheme: os.Getenv("URL_SCHEME"),
		User:   url.UserPassword(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD")),
		Host:   "localhost:5432",
		Path:   os.Getenv("POSTGRES_DB"),
	}

	urlValues := dsn.Query()                       //Query parses RawQuery and returns the corresponding values
	urlValues.Add(os.Getenv("SSLMODE"), "disable") // disbale sslmode and add to the map of Values

	// database url = URL_SCHEME://POSTGRES_USER:POSTGRES_PASSWORD@HOST:PORT/POSTGRES_DB?sslmode=SSLMODE
	dsn.RawQuery = urlValues.Encode()

	// create db connection
	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("sql open", err)
	}

	return db
}
