package main

import (
	"fmt"
	"time"
)

// consumer
func main() {

	stopCh := make(chan struct{})
	ch := genNos(stopCh)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		close(stopCh)
		// stopCh <- struct{}{}
	}()
	for val := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(val)
	}
	fmt.Println("Done")
}

// producer (produce values for 5 seconds)
func genNos(stopCh chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for idx := 1; ; idx++ {
			select {
			case <-stopCh:
				fmt.Println("stop signal received!")
				break LOOP
			case ch <- 10 * idx:
			}
		}
		close(ch)
	}()
	return ch
}

/*
func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
*/
