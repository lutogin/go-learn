package functions

import (
	"fmt"
	"go-learn/basic/common"
)

func Run() {
	fmt.Println("Функции с реструктуризацией параметров")
	fmt.Println(common.Add(123, 2, 7))
	fmt.Println(common.Add([]int{123, 2, 7}...))

	fmt.Println("Функция с возвратом нескольких значений")
	fibo1, fibo2 := common.FiboTest(1, 2)
	fmt.Println(fibo1, fibo2)

	fmt.Println("Функция как параметр")
	fmt.Println(common.Action(11, 22, common.SimpleAdd))

	fmt.Println("Анонимная функция")
	anonymusFn := func(a int, b int) int {
		return a + b
	}

	fmt.Println(anonymusFn(1, 2))

	fmt.Println("Замыкание")
	f := common.Square()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("Рекурсия (факториал + фибо)")
	fmt.Println(common.Factorial(5))
	fmt.Println(common.Fibo(8))

	fmt.Println("Panic test")
	fmt.Println(common.Devide(2, 3))
	//    fmt.Println(devide(3, 0))
}
