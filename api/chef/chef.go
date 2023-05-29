package chef

import (
	"net/http"
	"encoding/json"
	"log"
)

type postBody struct {
	code int `json:code`
}

func PostMethod(rw http.ResponseWriter, req *http.Request) int {

	var post postBody 

	if req.Method == "POST" {
		err := json.NewDecoder(req.Body).Decode(&post)

		if err != nil {
			log.Fatal("error decoding into int")
		}
		return post.code
	} else {
		return 0
	}
}