package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 200; i++ {
		// execute the increment() function concurrently
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(counter.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
