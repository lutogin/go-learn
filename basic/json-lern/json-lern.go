package json_lern

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string `json:"name"`
	Age       uint   `json:"age"`
	IsBlocked bool   `json:"isBlocked"`
}

func Run() {
	//sv := []string{"a", "b", "c"}
	//sv2 := map[string]interface{}{"field1": 1, "field2": "text", "field3": true}

	user := User{"Max", 20, true}

	setOfBytes, _ := json.Marshal(user) // serialize

	fmt.Printf("Marshal raw: %s \n", setOfBytes)
	fmt.Printf("Marshal string: %s \n", string(setOfBytes))

	container := User{}

	if err := json.Unmarshal(setOfBytes, &container); err != nil { // deserialize
		panic(err)
	}
	fmt.Printf("Unmarshal data: name = %s, age: %d \n", container.Name, container.Age)

	container2 := map[string]interface{}{}

	if err := json.Unmarshal(setOfBytes, &container2); err != nil { // deserialize2
		panic(err)
	}

	fmt.Printf("Unmarshal data to Set: name = %s, age: %d \n",
		container2["name"].(string),
		uint(container2["age"].(float64)),
	) // container["name"].(string) - casting data!!!
}
