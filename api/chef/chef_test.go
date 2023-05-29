package chef

import (
	"testing"
	"restaurantapi/webserver"
	"net/http"
)

func TestChef(t *testing.T) {

	http.HandleFunc("/chef", func(rw http.ResponseWriter, req *http.Request) {
		TestPostMethod(t)
	})

	// directory of web files
	webserver.WebFilesDir = "../../static"

	// start server
	webserver.RunServer(webserver.WebFilesDir)

}

func TestPostMethod(t *testing.T) {
	 
	var rw http.ResponseWriter
	var req *http.Request

	code := PostMethod(rw, req)
	expected := 200

	if code != expected {
		t.Errorf("expected %v, got %v", expected, code)
	}
}

