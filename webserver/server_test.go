// Testing Handlers

package webserver

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"restaurantapi/api/chef"
	"restaurantapi/api/menu"
	"testing"
)

func TestFileServer(t *testing.T) {

	t.Parallel()

	// returns a handler that serves HTTP request with file contents
	fileHandler := http.FileServer(http.Dir("../static"))

	//starts and returns a test server
	//automatically chooses a port to connect to
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

	// test body of response
	if body == nil {
		t.Errorf("want nil, got %v", string(body))
	}

}

func TestHandlers(t *testing.T) {

	t.Parallel()

	testHandlers := func(t testing.TB, handlerfunc func(rw http.ResponseWriter, req *http.Request), endpoint string, code int) {

		t.Helper() // prints the line number of the function call (testHandlers)

		// captures everything that is written with the ResponseWriter and returns ResponseRecorder
		rec := httptest.NewRecorder()

		// returns an incoming server request
		req := httptest.NewRequest(http.MethodGet, endpoint, nil)
		// if err != nil {
		// 	t.Error(err)
		// }

		handlerfunc(rec, req) // call endpoint handler

		resp := rec.Result() // returns the response generated by the handler

		if resp.StatusCode != code {
			t.Errorf("want %d, got %d", code, rec.Result().StatusCode)
		}

		// type A struct {
		// 	data string
		// }
		// a := A{data: ""}
		// defer rec.Result().Body.Close()
		// errBody := json.NewDecoder(rec.Result().Body).Decode(&a)
		// // get the response body
		// if errBody != nil {
		// 	t.Error(errBody)
		// }

		// if body != nil {
		// 	t.Errorf("want nil, got %v", string(body))
		// }

	}

	testCases := []struct {
		handlerName string
		handlerfunc func(rw http.ResponseWriter, req *http.Request)
		endpoint    string
		statusCode  int
	}{
		{handlerName: "ChefHandler", handlerfunc: chef.NewChef().ChefHandler, endpoint: "/chef", statusCode: http.StatusOK},
		{handlerName: "MenuHandler", handlerfunc: menu.NewMenu().MenuHandler, endpoint: "/menu", statusCode: http.StatusOK},
	}

	for _, testCase := range testCases {

		testCase := testCase

		// run subtests
		t.Run(testCase.handlerName, func(t *testing.T) {

			testHandlers(t, testCase.handlerfunc, testCase.endpoint, testCase.statusCode)

		})

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
