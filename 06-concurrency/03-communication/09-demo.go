package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory (not advisable)
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, 100, 200)
	wg.Wait()
	fmt.Println("result :", result)
}

func add(wg *sync.WaitGroup, x, y int) {
	defer wg.Done()
	result = x + y
}
