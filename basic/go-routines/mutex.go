package go_routines

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock() // Lock access to the common field (counter) for change only once per time
	c.counter[key]++
	c.mu.Unlock() // Unlock
}

func (c *Counter) Get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter[key]
}

func RunMutex() {
	key := "test"
	counter := Counter{counter: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go counter.Inc(key)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Current value: %d", counter.Get(key))
}
