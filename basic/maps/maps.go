package maps

import "fmt"

func Run() {
	fmt.Println("Map")
	var map1 = map[string]string{
		"key1": "val1",
		"key2": "val2",
	}
	fmt.Println(map1)

	map2 := make(map[string]string)

	map2["a1"] = "val1"
	map2["a2"] = "val2"

	for _, v := range map2 {
		fmt.Println(v)
	}

	fmt.Println("map2 len = ", len(map2))

	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}

	if val, ok := people["Tom"]; ok {
		fmt.Println("Checked key at MAP, result: ", ok, "value: ", val)
	}

	people["AddNewUser"] = 22 // adding new user

	delete(people, "Tom") // remove

	if _, ok := people["Tom"]; ok {
		fmt.Println("Checked key at MAP after deleting, result: ", ok)
	} else {
		fmt.Println("Key is deleted.")
	}

	mapWithAny := make(map[string]interface{})
	mapWithAny["bool"] = true
	mapWithAny["str"] = "string"
	mapWithAny["num"] = 12

	fmt.Println(mapWithAny)
}