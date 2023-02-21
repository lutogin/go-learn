package methods

import "fmt"

type library []string

// func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращаемых_результатов){
//     тело_метода
// }

func (l library) print() {

	for _, val := range l {
		fmt.Println(val)
	}
}

var lib library = library{"Book1", "Book2", "Book3"}

// методы структур

type person struct {
	age  uint
	name string
}

func (p person) PrintMe() {
	fmt.Println("Age is: ", p.age, "Name is: ", p.name)
}

var p1 person = person{21, "Andre"}

func Run() {
	lib.print()
	p1.PrintMe()
}
