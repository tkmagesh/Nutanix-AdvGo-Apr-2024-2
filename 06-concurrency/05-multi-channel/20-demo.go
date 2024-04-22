package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for val := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(val)
	}
	fmt.Println("Done")
}

// producer (produce values for 5 seconds)
func genNos(ch chan int) {
	timeoutCh := timeout(5 * time.Second)
LOOP:
	for idx := 1; ; idx++ {
		select {
		case <-timeoutCh:
			fmt.Println("timeout occurred!")
			break LOOP
		case ch <- 10 * idx:
		}
	}
	close(ch)
}

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
