package codintions

import "fmt"

func test() (int, int) {
	return 1, 2
}

func Run() {
	fmt.Println("Ифчик")

	if 2 > 3 || 3 < 5 {
		fmt.Println("Check the IF")
	}

	if first, _ := test(); first == 1 {
		fmt.Println("Possible way to use IF")
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
