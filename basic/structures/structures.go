package structures

import "fmt"

func Run() {
	type person struct {
		name string
		age  int
	}

	var user1 person = person{"Nick", 21}
	var user1Pointer *person = &user1

	fmt.Println("age: ", user1.age, "name: ", user1.name, "pointer: ", user1Pointer)

	// nested struct

	type contact struct {
		email string
		phone string
	}

	type nperson struct {
		name        string
		age         int
		contactInfo contact
	}

	var nUser1 nperson = nperson{"Tim", 21, contact{"test@gmail.com", "123"}}
	var nUser2 nperson = nperson{name: "John", age: 25, contactInfo: contact{email: "test@gmail.com", phone: "567"}}
	fmt.Println(nUser1, nUser2)

	var nUser3Link nperson = nUser1

	nUser3Link.age = 23

	fmt.Println("is it copy?: ", nUser3Link)
	fmt.Println("original: ", nUser1)

	asd := 1
	var fgh uint = uint(asd)
	asd = 2
	fmt.Println(asd, fgh)
}
