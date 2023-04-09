package json_lern

import (
	"fmt"
	"github.com/tidwall/gjson"
)

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type UserDB struct {
	Name Name
	Age  uint `json:"age"`
}

func RunGson() {
	// gjson
	//user := UserDB{Name: Name{First: "Stat", Last: "Copilov"}, Age: 39}
	user2 := `{"name": {"first": "Stas", "last": "Copilov"}, "age": 39}`

	name := gjson.Get(user2, "name.first") // like lodash.get() - getting data from JSON

	fmt.Printf("Parsed name: %s", name)
}
