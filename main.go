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
		log.Fatal("Error loading .env file")
	}

	utils.Database = database.Conn()                                 // establish db connections and return the database
	utils.CreateTables("database/create_tables.sql", utils.Database) // read sql script and create db tables

}
