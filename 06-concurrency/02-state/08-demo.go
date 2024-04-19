// using sync/atomic package

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

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
	atomic.AddInt64(&counter, 1)

}
