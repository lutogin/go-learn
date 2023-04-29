package io_console

import (
	"bufio"
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

	fmt.Print("Enter a expression splits by `,`: ") // for getting all string
	in := bufio.NewReader(os.Stdin)                 // read all string from console
	exps, err := in.ReadString('\n')

	// Remove the newline character at the end of the input
	exps = exps[:len(exps)-1]

	fmt.Println(exps, err)
}
