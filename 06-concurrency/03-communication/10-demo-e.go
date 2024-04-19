package main

import (
	"fmt"
	"time"
)

// share memory by communicating

func main() {
	ch := add(100, 200)
	/* go func() {
		ch <- 10000
	}() */
	result := <-ch
	fmt.Println("result :", result)
}

func add(x, y int) <-chan int /* receive only channel */ {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
