package go_routines

import (
	"fmt"
	"sync"
	"time"
)

type CounterRW struct {
	mu      sync.RWMutex
	counter map[string]int
}

func (c *CounterRW) CountMe(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter[key]
}

func (c *CounterRW) CountMeAgain(key string) int {
	c.mu.RLock() // don't block if there aren't block to record Use RWLock if critic code don't change a data, just read
	defer c.mu.Unlock()
	return c.counter[key]
}

func RunRWMutex() {
	key := "test"
	counter := Counter{counter: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go counter.Inc(key)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Current value: %d", counter.Get(key))
}
