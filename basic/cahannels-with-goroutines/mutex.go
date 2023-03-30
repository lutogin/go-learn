package cahannels_with_goroutines

import (
	"fmt"
	"sync"
)

/**
Мьютексы позволяют разграничить доступ к некоторым общим ресурсам, гарантируя,
что только одна горутина имеет к ним доступ в определенный момент времени.
И пока одна горутина не освободит общий ресурс, другая горутина не может с ним работать.

С помощью мьютексов можно ограничить доступ к переменной таким образом,
чтобы только одна горутина имела к ней монопольный доступ в один момент времени
*/

var counter int = 0 // common resource

func RunWithoutMutex() {
	ch := make(chan bool)

	for i := 1; i < 5; i++ {
		go work(i, ch) // run 5 go routines with which change common variable "counter"
	}

	// ожидаем завершения всех горутин
	for i := 1; i < 5; i++ {
		<-ch
	}
	fmt.Println("The End")
}

func work(number int, ch chan<- bool) {
	counter = 0

	for k := 1; k < 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}

	ch <- true
}

func RunMutex() {
	ch := make(chan bool)
	var mutex sync.Mutex

	for i := 1; i < 5; i++ {
		go workMutex(i, ch, &mutex) // run 5 go routines with which change common variable "counter"
	}

	// ожидаем завершения всех горутин
	for i := 1; i < 5; i++ {
		<-ch
	}
	fmt.Println("The End.")
}

func workMutex(number int, ch chan bool, mutex *sync.Mutex) {
	//defer mutex.Unlock() // unlock the common variable
	mutex.Lock() // lock the common variable
	counter = 0

	for k := 1; k < 5; k++ {
		counter++
		fmt.Println("Goroutine with mutex", number, "-", counter)
	}

	mutex.Unlock() // unlock the common variable
	ch <- true
}
