package go_routines

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(exp string, ch chan string) {
	splitData := strings.Split(exp, " ")

	if len(splitData) != 3 {
		ch <- errors.New("string must contain three basic math operations, like '3 + 2'").Error()
	}

	x, err := strconv.ParseInt(splitData[0], 10, 0)
	if err != nil {
		ch <- err.Error()
	}

	y, err := strconv.ParseInt(splitData[2], 10, 0)
	if err != nil {
		ch <- err.Error()
	}

	switch splitData[1] {
	case "+":
		ch <- fmt.Sprintf("%s = %d", exp, x+y)
	case "-":
		ch <- fmt.Sprintf("%s = %d", exp, x-y)
	case "*":
		ch <- fmt.Sprintf("%s = %d", exp, x*y)
	case "/":
		ch <- fmt.Sprintf("%s = %d", exp, x/y)
	default:
		ch <- errors.New(fmt.Sprintf("%s wrong math operand, available values '-, +, *, /'", exp)).Error()
	}
}

func RunAsyncCalc() {
	fmt.Print("Enter a expression splits by `,`: ")
	in := bufio.NewReader(os.Stdin) // read all string from console
	exps, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Error with entered data.")
	}

	expsChunk := strings.Split(exps, ",")
	ch := make(chan string)
	defer close(ch)

	for _, exp := range expsChunk {
		go calculate(strings.TrimSpace(exp), ch)
	}

	for i := 0; i < len(expsChunk); i++ {
		println(<-ch)
	}

	fmt.Println("Done.")
}
