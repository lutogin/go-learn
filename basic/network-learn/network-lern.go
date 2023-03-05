package network_learn

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RunRequest() {
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: go.dev\n\n"

	conn, err := net.Dial("tcp", "go.dev:80")

	if err != nil {
		fmt.Printf("Error during request. %s", err)
		return
	}

	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done.")
}

func RunListener() {
	message := "Hi I'm a server."

	listener, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()
	fmt.Println("Server is running...")
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		conn.Write([]byte(message))
		conn.Close()
	}
}
