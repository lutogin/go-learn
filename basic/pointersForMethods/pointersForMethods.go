package pointersForMethods

import "fmt"

type person struct {
	age  uint
	name string
}

func (p *person) setAge(newAge uint) {
	p.age = newAge
}

func Run() {
	var p1 person = person{21, "Nick"}

	fmt.Println("Before: ", p1.age)

	p1.setAge(33)

	fmt.Println("After: ", p1.age)
}
