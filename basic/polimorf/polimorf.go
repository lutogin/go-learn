package polimorf

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {
	model string
}

type Aircraft struct {
	model string
}

func (c *Car) move() {
	fmt.Println(c.model, " driving.")
}

func (a *Aircraft) move() {
	fmt.Println(a.model, " flying.")
}

func Run() {
	c1 := Car{"bmw"}
	c2 := Car{"mers"}
	a1 := Aircraft{"boeing"}

	var vehicles []Vehicle = []Vehicle{&c1, &c2, &a1}

	for _, val := range vehicles {
		val.move()
	}
}
