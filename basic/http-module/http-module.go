package http_module

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func RunGet() {
	res, err := http.Get("https://google.com/")

	if err != nil {
		println(err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func RunCustomClient() {
	client := http.Client{
		Timeout: time.Second * 6,
	}

	res, err := client.Get("https://google.com")

	if err != nil {
		println(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func RunDoRequest() {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://google.com", nil,
	)
	// добавляем заголовки
	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
