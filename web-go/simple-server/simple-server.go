package simple_server

import (
	"fmt"
	"net/http"
)

type handler struct {
	msg string
}

func (m handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.msg)
}

func RunSimpleServer() {
	fmt.Println("Server is running...")
	messageHandler := handler{"Hello world!"}

	err := http.ListenAndServe("127.0.0.1:8181", messageHandler)

	if err != nil {
		fmt.Printf("Server is die. %s", err.Error())
	}
}

func RunServerWithRouts() {
	http.Handle("/about", &handler{"About."})
	http.Handle("/contact", &handler{"Contact page."})
	http.Handle("/", &handler{"Index page."})

	err := http.ListenAndServe("127.0.0.1:8181", nil)
	if err != nil {
		fmt.Printf("Server is die. %s", err.Error())
	} else {
		fmt.Println("Server is running...")
	}
}
