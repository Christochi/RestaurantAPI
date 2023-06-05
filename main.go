package main

import (
	"restaurantapi/webserver"
)

func main() {

	// directory of web files
	webserver.WebFilesDir = "./static"

	webserver.RunServer(webserver.WebFilesDir)

}
