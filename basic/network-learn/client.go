package network_learn

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:4545")

	if err != nil {
		fmt.Printf("Error durung establishing the connection. %s", err)
		return
	}

	conn.SetReadDeadline(time.Now().Add(time.Second * 5)) // 5 sec timeout

	defer conn.Close()

	for {
		var word string

		fmt.Print("Enter a word: ")
		_, errEnter := fmt.Fscan(os.Stdin, &word) // Fscan or Scanln or Scanf

		if errEnter != nil {
			fmt.Printf("Error on type. %e", errEnter)
			continue
		}

		if n, errWrite := conn.Write([]byte(word)); n == 0 || errWrite != nil {
			fmt.Printf("Enter during send. %s", errWrite)
			return
		}

		buff := make([]byte, 1024)
		n, errRead := conn.Read(buff)

		if errRead != nil {
			fmt.Printf("Enter during receive. %s", errRead)
			break
		}

		fmt.Print("Translation: ")
		fmt.Println(string(buff[0:n]))
		fmt.Println()
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700)) // 700 ms timeout for receive data
	}
}
