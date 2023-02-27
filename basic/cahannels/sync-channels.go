package cahannels

import (
	"fmt"
	"go-learn/basic/common"
)

func RunSync() {
	intCh := make(chan int)

	go common.FactorialChannel(5, intCh)

	fmt.Println("Factorial from goroutine: ", <-intCh)

	// ---------------------------
	fmt.Println("---------------------------")

	structCh := make(chan struct{})
	results := make(map[int]int)

	go common.FactorialWithStruct(6, structCh, results)

	//<-structCh // wait the closing the channel. It's a example like async await in JS

	for i, v := range results {
		fmt.Println(i, " - ", v)
	}
}
