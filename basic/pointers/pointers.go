package pointers

import "fmt"

func Run() {
	var p *int
	n := 10
	p = &n

	fmt.Println("n value: ", n, "n address: ", p, "value by pointer: ", *p)

	*p = 99 // update value by pointer

	f := 2.3
	pf := &f

	fmt.Println("Address:", pf)
	fmt.Println("Value:", *pf)

	// pointers in function's params
	var zxc uint = 5

	testPointer := func(x *uint) {
		*x = (*x) * (*x)
	}

	testPointer(&zxc)

	fmt.Println(zxc)
}
