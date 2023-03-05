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
	listener, err := net.Listen("tcp4", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()
	fmt.Println("Server is running...")

	message := "Hi I'm a server."

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

func RunServer() {
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	fmt.Println("The server is running...")

	for {
		conn, errConn := listener.Accept()

		fmt.Println("Accepted data.")

		if errConn != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}

		go handleConnection(conn)
	}
}

var dict = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
}

// обработка подключения
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		fmt.Println("Handled in connection.")
		// считываем полученные в запросе данные
		input := make([]byte, 4096)
		n, err := conn.Read(input)

		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		source := string(input[0:n])
		fmt.Printf("SORCE: %s", source)
		// на основании полученных данных получаем из словаря перевод
		target, ok := dict[source]
		if ok == false { // если данные не найдены в словаре
			target = "undefined"
		}
		// выводим на консоль сервера диагностическую информацию
		fmt.Println(source, "-", target)
		// отправляем данные клиенту
		conn.Write([]byte(target))
	}
}
