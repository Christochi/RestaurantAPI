package chef

import (
	"testing"
	"restaurantapi/webserver"
)

func TestChef(t *testing.T) {

	// directory of web files
	webserver.WebFilesDir = "../../static"
	
	// start server
	webserver.RunServer(webserver.WebFilesDir)

}

