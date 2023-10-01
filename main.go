package main

import (
	"log"
	"restaurantapi/database"
	"restaurantapi/utils"
	"restaurantapi/webserver"

	"github.com/joho/godotenv"
)

func main() {

	// directory of web files
	webserver.WebFilesDir = "./static"

	// connect to database and create db tables
	dbConn()

	// start the server
	webserver.RunServer(webserver.WebFilesDir)

}

func dbConn() {

	// load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Conn()                // establish db connections and return the database
	utils.OpenFile("chef_table.sql", db) // read and create SQL table
}
