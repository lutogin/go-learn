package go_routines

import (
	"fmt"
	"time"
)

func RunSelect() {
	chData := make(chan int)
	chExit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-chData)
		}
		chExit <- true
	}()

	selectOne(chData, chExit)
}

func selectOne(data chan int, exit chan bool) {
	i := 0
	for {
		select {
		case data <- i:
			i += 1
		case <-exit:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Waiting...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
