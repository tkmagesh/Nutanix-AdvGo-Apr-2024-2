package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 200; i++ {
		// execute the increment() function concurrently
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(counter)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()
}
