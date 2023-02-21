package vars

import (
    "fmt"
    "math"
)

func Run() {
    fmt.Println("Переменные")
    const pi float32 = 3.14
    az := 1
    z := math.Round(float64(pi + float32(az)))
    fmt.Println(z)
}