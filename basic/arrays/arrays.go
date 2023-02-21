package arrays

import "fmt"

func Run() {
    fmt.Println("Массивы")
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)

    arr2 := [...]string{"a", "b"}
    fmt.Println(arr2)
    fmt.Println("Len of array is", len(arr2))
    arr3 := []string{"a", "b", "c"} // slice
    fmt.Println("Len of slice is", len(arr3))

    fmt.Println("Перебор массива")
    users := []string{"alice", "bob", "jack"}

    for index, value := range users {
        fmt.Println(index, value)
    }

    for _, value := range users {
        fmt.Println(value)
    }
}