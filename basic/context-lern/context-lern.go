package context_lern

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Run() {
	ctx := context.Background() // return new context
	ctx = context.WithValue(ctx, "key", "hello")
	ctx, cancel := context.WithCancel(ctx) // make new ctx and possible to cancel it manually

	//ctx := context.WithDeadline(ctx, time.Now().Add(1*time.Second)) // close context by the time
	//ctx := context.WithTimeout(ctx, 1*time.Second) // close context by timer

	go func() {
		err := cancelRequest(ctx)

		if err != nil {
			cancel()
		}
	}()

	makeRequest(ctx, "https://google.com")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(5000 * time.Millisecond)

	return fmt.Errorf("Filed request.")
}

func makeRequest(ctx context.Context, reqStr string) {
	req, _ := http.NewRequest(http.MethodGet, reqStr, nil)
	req = req.WithContext(ctx)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Response completed, status: %d \n", res.StatusCode)
		fmt.Printf("Context value: %s", ctx.Value("key"))
	case <-ctx.Done(): // handle when channel will be close
		fmt.Println("Request failed. Too long.")
	}
}
