package chef

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

// Data about a chef
type Chef struct {

	Name string `json:"name"`
	About string `json:"about"`
}

var chef []Chef // list of chefs 

// client send chef data using POST Method
func PostMethod(rw http.ResponseWriter, req *http.Request) {

	// decode json to struct
	err := json.NewDecoder(req.Body).Decode(&chef)

	// error handling
	if err != nil {
		log.Fatal("error decoding into struct")
	}

	fmt.Printf("%+v\n", chef)
	
	// server's response to client
	fmt.Fprintln(rw, "success")	
}

// client requests for chef data using GET Method
func GetMethod(rw http.ResponseWriter, req *http.Request) {



}