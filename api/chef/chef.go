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

func PostMethod(rw http.ResponseWriter, req *http.Request) {

	var p []PostBody 

	err := json.NewDecoder(req.Body).Decode(&p)

	if err != nil {
		log.Fatal("error decoding into struct")
	}

	fmt.Printf("%+v\n", p)
	
	fmt.Fprintln(rw, "success")
		
	
}