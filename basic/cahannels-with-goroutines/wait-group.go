package cahannels_with_goroutines

import (
	"fmt"
	"sync"
	"time"
)

/*
Этот тип позволяет определить группу горутин, которые должны выполняться вместе как одна группа.
И можно установить блокировку, которая приостановит выполнение функции, пока не завершит выполнение вся группа горутин
*/

func RunWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2) // группа из 2 готрутин

	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}

	work(1) // вызываем горутины
	work(2)

	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")
}
