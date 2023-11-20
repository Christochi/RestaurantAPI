package menu

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostMenu(t *testing.T) {

	t.Parallel()

	menu := NewMenu()      // chef object
	b := new(bytes.Buffer) // zero value of buffer of bytes

	// test data
	testData := []menuJson{
		{
			Type:      "Breakfast",
			Meal:      "Toasted Bread with Mashed Potato and Fried Egg",
			Price:     "$10",
			Desc:      "Mildly tasty mashed potato and egg plastered on toasted bread",
			Image:     "toasted-bread-potato",
			Available: true,
		},
		{
			Type:      "Lunch",
			Meal:      "Fried Rice with Salmon",
			Price:     "$15",
			Desc:      "Fried rice with baked chicken",
			Image:     "fried-rice-chicken",
			Available: true,
		},
		{
			Type:      "Dinner",
			Meal:      "Spinach Beef",
			Price:     "$25",
			Desc:      "spinach-beef",
			Image:     "spinach-beef",
			Available: true,
		},
	}

	// encode to bytes (buffer of bytes implements io.Writer interface)
	_ = json.NewEncoder(b).Encode(testData)

	// captures everything that is written with the ResponseWriter and returns ResponseRecorder
	rec := httptest.NewRecorder()

	// creates a request (buffer of bytes implements io.Reader interface)
	req := httptest.NewRequest(http.MethodPost, "/menu", b)

	menu.MenuHandler(rec, req) // call postMenu

	// compares menu value and testData
	for i, c := range *menu {
		for k, data := range testData {

			if i == k && c != data {
				t.Errorf("want %#v\n, got  %#v\n", c, data)
			}
		}
	}
}
