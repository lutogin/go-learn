package cahannels_with_goroutines

import (
	"fmt"
	"go-learn/basic/common"
)

func Run() {
	// -------------------- non buffer channels
	//var intCh chan int
	var intCh chan int = make(chan int)
	go func() {
		fmt.Println("Go routine starts.")
		intCh <- 5 // blocks due main get data from the channel
	}()

	fmt.Println(<-intCh) // blocks due main get data from the channel
	fmt.Println("End.")

	// ---------------- 2
	var intCh2 chan int = make(chan int)

	go common.FactorialChannel(5, intCh2) // pass the channel

	fmt.Println(<-intCh2) // read from channel

	// -------------------- buffer channels
	var chInt3 chan int = make(chan int, 3)

	chInt3 <- 2
	chInt3 <- 5
	fmt.Println("cap(chInt3) = ", cap(chInt3)) // 2 (current size)
	fmt.Println("len(chInt3) = ", len(chInt3)) // 3 length
	chInt3 <- 8

	fmt.Println(<-chInt3)
	fmt.Println(<-chInt3)
	fmt.Println(<-chInt3)

	close(chInt3) // close the channel!!! Important for streaming

	if val, opened := <-chInt3; opened { // possible IF
		fmt.Println(val)
	} else {
		fmt.Println("Channel closed!")
	}

	//var inCh chan<- int  // channel only for sending
	//var outCh <-chan int // channel only for receiving

	// chan as returned argument
	fmt.Println("chan as returned argument: ")
	fmt.Println(<-createChan(50))
}

func createChan(n int) <-chan int {
	ch := make(chan int) // создаем канал
	go func() {
		ch <- n // отправляем данные в канал ТОЛЬКО через запущеную горутину
	}()

	return ch // возвращаем канал
}
