package io_learn

import "fmt"

type phoneNumber struct{}

func (ph phoneNumber) Write(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	for i := 0; i < len(data); i++ {
		if data[i] >= '0' && data[i] <= '9' {
			fmt.Print(string(data[i]))
		}
	}

	fmt.Println()
	return len(data), nil
}

func RunWrite() {
	bytes1 := []byte("+1(234)567 9010")
	bytes2 := []byte("+2-345-678-12-35")

	writer := phoneNumber{}

	writer.Write(bytes1)
	writer.Write(bytes2)
}
