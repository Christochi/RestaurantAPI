// Testing Handlers

package webserver

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"restaurantapi/api/chef"
	"restaurantapi/api/menu"
	"restaurantapi/utils"
	"testing"
)

func TestNoCache(t *testing.T) {

	t.Parallel()

	// returns a handler that serves HTTP request with file contents
	fileHandler := http.FileServer(http.Dir("../static"))

	//starts and returns a test server
	//automatically chooses an available port to connect to
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

		// run subtests
		t.Run(key, func(t *testing.T) {

			actual := resp.Header.Values(key) // return slice of values associated with the key

			if !reflect.DeepEqual(actual, expectedValue) {
				t.Errorf("want %v, got %v", expectedValue, actual)
			}

		})

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

	testHandlers := func(t testing.TB, handlerfunc func(rw http.ResponseWriter, req *http.Request), endpoint string, code int, jsonData any) {

		// captures everything that is written with the ResponseWriter and returns ResponseRecorder
		rec := httptest.NewRecorder()

		// returns an incoming server request
		req := httptest.NewRequest(http.MethodGet, endpoint, nil)

		handlerfunc(rec, req) // call endpoint handler

		resp := rec.Result() // returns the response generated by the handler

		if resp.StatusCode != code {
			t.Errorf("want %d, got %d", code, rec.Result().StatusCode)
		}

		defer resp.Body.Close()

		// read and decode Response's body to struct
		utils.Post(rec, req, jsonData)

		// test if the object is not nil
		if !reflect.Indirect(reflect.ValueOf(jsonData)).IsNil() {
			t.Errorf("want nil, got %#v", jsonData)
		}

	}

	// objects from api package
	chef := chef.NewChef()
	menu := menu.NewMenu()

	testCases := []struct {
		handlerName string
		handlerfunc func(rw http.ResponseWriter, req *http.Request)
		endpoint    string
		statusCode  int
		jsonData    any
	}{
		{handlerName: "ChefHandler", handlerfunc: chef.ChefHandler, endpoint: "/chef", statusCode: http.StatusOK, jsonData: chef},
		{handlerName: "MenuHandler", handlerfunc: menu.MenuHandler, endpoint: "/menu", statusCode: http.StatusOK, jsonData: menu},
	}

	for _, testCase := range testCases {

		testCase := testCase

		// run subtests
		t.Run(testCase.handlerName, func(t *testing.T) {

			testHandlers(t, testCase.handlerfunc, testCase.endpoint, testCase.statusCode, testCase.jsonData)

		})

	}
}
