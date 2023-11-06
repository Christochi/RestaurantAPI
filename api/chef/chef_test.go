// Testing Chef API

package chef

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test data is decoded to struct
func TestPostChef(t *testing.T) {

	t.Parallel()

	chef := NewChef()      // chef object
	b := new(bytes.Buffer) // zero value of buffer of bytes

	// test data
	testData := []chefJson{
		{
			Name:   "Chocho Okoye",
			About:  "Assistant Chef",
			Image:  "chocho.jpg",
			Gender: "M",
			Age:    2,
		},
		{
			Name:   "John Abraham",
			About:  "John Abraham has 7 years experience making delicious meals for 5-star hotels and famous restaurants in North America",
			Image:  "johnabraham.jpg",
			Gender: "M",
			Age:    35,
		},
		{
			Name:   "John Doe",
			About:  "John Doe has 20 years experience cooking for famous restaurants in African and the Caribbean",
			Image:  "johndoe.jpg",
			Gender: "M",
			Age:    50,
		},
	}

	// encode to bytes, buffer of bytes implements io.Writer interface
	_ = json.NewEncoder(b).Encode(testData)

	// captures everything that is written with the ResponseWriter and returns ResponseRecorder
	rec := httptest.NewRecorder()

	// creates a request
	req := httptest.NewRequest(http.MethodPost, "/chef", b)

	chef.ChefHandler(rec, req) // call postChef

	for i, c := range *chef {
		for k, data := range testData {

			if i == k && c != data {
				t.Errorf("want %#v\n, got  %#v\n", c, data)
			}
		}
	}

}
