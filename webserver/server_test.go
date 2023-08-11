// Testing Handlers
package webserver

import (
	"net/http"
	"net/http/httptest"
	"restaurantapi/api/chef"
	"restaurantapi/api/menu"
	"testing"
)

func TestChefHandler(t *testing.T) {

	t.Parallel()

	newChef := chef.NewChef() // chef object

	//starts and returns a server
	server := httptest.NewServer(http.HandlerFunc(newChef.ChefHandler))

	url := server.URL + "/chef" // url

	// creates a request to query the server
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error(err)
	}

	// send the HTTP request and return an HTTP response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	statusCode := resp.StatusCode // return status code gotten after hitting the endpoint
	want := 200

	if statusCode != want {
		t.Errorf("want %d, got %d", want, statusCode)
	}

	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body) // get the response body
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if body != nil {
	// 	t.Errorf("want nil, got %v", body)
	// }

}

func TestMenuHandler(t *testing.T) {

	t.Parallel()

	newMenu := menu.NewMenu() // menu object

	//starts and returns a server
	server := httptest.NewServer(http.HandlerFunc(newMenu.MenuHandler))

	url := server.URL + "/menu" // url

	// creates a request to query the server
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error(err)
	}

	// send the HTTP request and return an HTTP response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	statusCode := resp.StatusCode // return status code gotten after hitting the endpoint
	want := 200

	if statusCode != want {
		t.Errorf("want %d, got %d", want, statusCode)
	}

	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body) // get the response body
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if body != nil {
	// 	t.Errorf("want nil, got %v", body)
	// }

}
