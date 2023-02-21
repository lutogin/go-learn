package slices

import "fmt"

func Run() {
	fmt.Println("Slice")
	var slice1 []string = make([]string, 3)
	slice1[0] = "bob"
	slice1[1] = "mike"
	slice1[2] = "alice"
	slice1 = append(slice1, "charley") // add to slice
	fmt.Println(slice1[0:2])
	fmt.Println(slice1[2:4])
	// remove from slice
	slice1 = append(slice1[:2], slice1[3:]...)
	fmt.Println(slice1)
}
