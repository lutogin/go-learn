package codintions

import "fmt"

func Run() {
    fmt.Println("Ифчик")

    if 2 > 3 || 3 < 5 {
        fmt.Println("Check the IF")
    }

    fmt.Println("Свитч")
    ab := 87
    switch ab {
        case 9:
            fmt.Println("a = 9")
        case 8:
            fmt.Println("a = 8")
        case 7:
            fmt.Println("a = 7")
        default:
            fmt.Println("значение переменной a не определено")
    }
}