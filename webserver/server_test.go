// Testing Handlers

package webserver

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFileServer(t *testing.T) {

	// returns a handler that serves HTTP request with file contents
	fileHandler := http.FileServer(http.Dir("../static"))

	//starts and returns a server
	server := httptest.NewServer(noCache(http.StripPrefix("/", fileHandler)))
	defer server.Close() // close server after all requests have been completed

	// sends a request to the URL
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	// test status code
	statusCode := resp.StatusCode
	if expectedStatusCode := http.StatusOK; statusCode != expectedStatusCode {
		t.Errorf("want %d, got %d", expectedStatusCode, statusCode)
	}

	// map of expected headers
	expectedHeader := map[string][]string{
		"pragma":                 {"no-cache"},
		"X-Content-Type-Options": {"nosniff"},
		"Cache-Control":          {"no-cache, no-store, must-revalidate"},
	}

	// test headers
	for key, expectedValue := range expectedHeader {

		actual := resp.Header.Values(key)

		if !reflect.DeepEqual(actual, expectedValue) {
			t.Errorf("want %v, got %v", expectedValue, actual)
		}

	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // get the response body
	if err != nil {
		t.Error(err)
	}

	if body == nil {
		t.Errorf("want nil, got %v", string(body))
	}

}

// func TestChefHandler(t *testing.T) {

// 	t.Parallel()

// 	newChef := chef.NewChef() // chef object

// 	//starts and returns a server
// 	server := httptest.NewServer(http.HandlerFunc(newChef.ChefHandler))

// 	url := server.URL + "/chef" // url

// 	// creates a request to query the server
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// send the HTTP request and return an HTTP response
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	statusCode := resp.StatusCode // return status code gotten after hitting the endpoint
// 	want := 200

// 	if statusCode != want {
// 		t.Errorf("want %d, got %d", want, statusCode)
// 	}

// 	// defer resp.Body.Close()
// 	// body, err := io.ReadAll(resp.Body) // get the response body
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }

// 	// if body != nil {
// 	// 	t.Errorf("want nil, got %v", body)
// 	// }

// }

// func TestMenuHandler(t *testing.T) {

// 	t.Parallel()

// 	newMenu := menu.NewMenu() // menu object

// 	//starts and returns a server
// 	server := httptest.NewServer(http.HandlerFunc(newMenu.MenuHandler))

// 	url := server.URL + "/menu" // url

// 	// creates a request to query the server
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// send the HTTP request and return an HTTP response
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	statusCode := resp.StatusCode // return status code gotten after hitting the endpoint
// 	want := 200

// 	if statusCode != want {
// 		t.Errorf("want %d, got %d", want, statusCode)
// 	}

// 	// defer resp.Body.Close()
// 	// body, err := io.ReadAll(resp.Body) // get the response body
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }

// 	// if body != nil {
// 	// 	t.Errorf("want nil, got %v", body)
// 	// }

// }
