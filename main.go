package main

import (
	"restaurantapi/webserver"
)

func main() {

	// directory of web files
	webserver.WebFilesDir = "./static"

	// start the server
	webserver.RunServer(webserver.WebFilesDir)

}
