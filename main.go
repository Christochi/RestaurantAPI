package main

import (
	"log"
	"restaurantapi/database"
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

	// load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file, ", err)
	}

	utils.Database = database.Conn()                           // confirm db uri and return the database
	query := utils.ReadSQLScript("database/create_tables.sql") // read sql script
	utils.ExecuteQueries(string(query), utils.Database)        // execute SQL queries

}
