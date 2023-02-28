package io_learn

import (
	"fmt"
	"io"
)

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0

	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}

	return count, io.EOF
}

func RunRead() {
	var ph1 phoneReader = "+1(234)567 9010"
	var ph2 phoneReader = "+2-345-678-12-35"

	buffer1 := make([]byte, len(ph1)) // create a buffer
	buffer2 := make([]byte, len(ph2))

	ph1.Read(buffer1)
	ph2.Read(buffer2)

	fmt.Println(string(buffer1))
	fmt.Println(string(buffer2))
}
