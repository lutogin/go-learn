package go_routines

import "fmt"

func RunWorkGoroutine() {
	ch1 := make(chan int)

	go goRoutine(ch1, 5)
	fmt.Printf("Factorial 5 is %d \n", <-ch1)

	// example with stream of data
	ch2 := make(chan int)

	go goRoutineWithStreamOfData(ch2, 6)

	for n := range ch2 {
		fmt.Printf("Iterate value from the stream: %d \n", n)
	}

	fmt.Println("Done.")
}

func goRoutine(ch chan int, n int) {
	result := n
	for i := 1; i < n; i++ {
		result *= i
	}

	ch <- result
}

func goRoutineWithStreamOfData(ch chan int, n int) {
	result := n
	for i := 1; i < n; i++ {
		result *= i
		ch <- result
	}
	close(ch) // CLOSE CHANNEL. Don't wait for me.
}
