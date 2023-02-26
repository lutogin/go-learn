package go_routines

import (
	"fmt"
	"go-learn/basic/common"
)

func Run() {
	fmt.Println("Start.")

	var i uint

	for i = 2; i <= 7; i++ {
		go func(n uint) {
			fmt.Println(common.Factorial(n))
		}(i)
	}

	fmt.Println("Press any key to continue.")

	fmt.Scanln() // wait

	fmt.Println("Finish.")
}
