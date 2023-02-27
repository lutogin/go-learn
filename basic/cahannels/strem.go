package cahannels

import (
	"fmt"
	"go-learn/basic/common"
)

func RunStream() {
	streamCh := make(chan int)

	go common.FactorialSteam(6, streamCh)

	//for {
	//	v, opened := <-streamCh // get data from the stream
	//	if !opened {
	//		break // if stream is close
	//	}
	//	fmt.Println(v)
	//}

	fmt.Println("-------------------------------")
	fmt.Println("other war for iterate")

	// Use FOR to count all elements of stream
	for v := range streamCh {
		fmt.Println(v)
	}
}
