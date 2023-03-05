package io_console

import (
	"fmt"
	"os"
)

func Run() {
	var name string
	var age uint

	fmt.Print("Enter name: ")
	fmt.Fscan(os.Stdin, &name) // Fscan or Scanln or Scanf
	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	fmt.Println(name, age)
}
