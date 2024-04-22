package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

// producer
func genNos(ch chan int) {
	ch <- 10
	time.Sleep(500 * time.Millisecond)
	ch <- 20
	time.Sleep(500 * time.Millisecond)
	ch <- 30
	time.Sleep(500 * time.Millisecond)
	ch <- 40
	time.Sleep(500 * time.Millisecond)
	ch <- 50
	time.Sleep(500 * time.Millisecond)
}
