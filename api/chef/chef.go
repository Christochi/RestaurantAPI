package chef

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

type PostBody struct {
	Code string `json:"code"`
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