package main

import (
	"log"
	"restaurantapi/database"
	errs "restaurantapi/errors"
	"restaurantapi/utils"
	"restaurantapi/webserver"

	"github.com/joho/godotenv"
)

func main() {

	// connect to database and create db tables
	dbConn()

	// directory of web files
	webserver.WebFilesDir = "./static"

	// start the server
	webserver.RunServer(webserver.WebFilesDir)

}

func dbConn() {

	var err error

	// load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file, ", err)
	}

	// confirm db uri and return the database
	utils.Database, err = database.Conn()

	// validate if a db driver was supplied
	if err != nil {
		log.Fatal(errs.DatabaseError(err))
	}

	// query, err := utils.ReadSQLScript("database/create_tables.sql") // read sql script
	// if err != nil {
	// 	errs.DatabaseError(err)
	// }

	// utils.ExecuteQueries(string(query), utils.Database) // execute SQL queries

}
