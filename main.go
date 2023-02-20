package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Переменные")
    const pi float32 = 3.14
    az := 1
    z := math.Round(float64(pi + float32(az)))
    fmt.Println(z)

    fmt.Println("Массивы")
    arr := [5]int{1,2,3,4,5}
    fmt.Println(arr)

    arr2 := [...]string{"a", "b"}
    fmt.Println(arr2)
    fmt.Println("Len of array is", len(arr2))
    arr3 := []string{"a", "b", "c"} // slice
    fmt.Println("Len of slice is", len(arr3))


    fmt.Println("Ифчик")

    if 2 > 3 || 3 < 5 {
        fmt.Println("Check the IF")
    }

    fmt.Println("Свитч")
    ab := 87
    switch(ab) {
        case 9:
            fmt.Println("a = 9")
        case 8:
            fmt.Println("a = 8")
        case 7:
            fmt.Println("a = 7")
        default:
            fmt.Println("значение переменной a не определено")
    }

    fmt.Println("Циклы")
    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }

    fmt.Println("Перебор массива")
    users := []string{"alice", "bob", "jack"}

    for index, value := range users {
        fmt.Println(index, value)
    }

    for _, value := range users {
        fmt.Println(value)
    }

    fmt.Println("Функции с реструктуризацией параметров")
    fmt.Println(add(123, 2, 7))
    fmt.Println(add([]int{123, 2, 7}...))

    fmt.Println("Функция с возвратом нескольких значений")
    fibo1, fibo2 := fiboTest(1, 2)
    fmt.Println(fibo1, fibo2)

    fmt.Println("Функция как параметр")
    fmt.Println(action(11, 22, simpleAdd))

    fmt.Println("Анонимная функция")
    anonymusFn := func(a int, b int) int {
        return a + b
    }

    fmt.Println(anonymusFn(1, 2))

    fmt.Println("Замыкание")
    f := square()
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())

    fmt.Println("Рекурсия (факториал + фибо)")
    fmt.Println(factorial(5))
    fmt.Println(fibo(8))

    fmt.Println("Panic test")
    fmt.Println(devide(2, 3))
//    fmt.Println(devide(3, 0))

    fmt.Println("Slice")
    var slice1 []string = make([]string, 3)
    slice1[0] = "bob"
    slice1[1] = "mike"
    slice1[2] = "alice"
    slice1 = append(slice1, "charley") // add to slice
    fmt.Println(slice1[0:2])
    fmt.Println(slice1[2:4])
    // remove from slice
    slice1 = append(slice1[:2], slice1[3:]...)
    fmt.Println(slice1)
}

func add(numbers ...int) int {
    sum := 0

    for _, value := range numbers {
        sum += value
    }

    return sum
}

func simpleAdd(a int, b int) int {
    return a + b
}

func fiboTest(init, final int) (int, int) {
    a := init + 1
    b := final + 2
    return a, b
}

func action(a int, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func square() func() int {
    var x int = 2
    return func() int {
        x++
        return x * x
    }
}

func factorial(n uint) uint {
    if n == 0 {
        return 1
    }

    return n * factorial(n - 1)
}

func fibo(n uint) uint {
    if n == 0 {
        return 0
    }

    if n == 1 {
        return 1
    }

    return fibo(n - 1) + fibo(n - 2)
}

func devide(a uint, b uint) float32 {
    if b < 1 {
        panic("Devide by zero impossible!")
    }

    return float32(a) / float32(b)
}