package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Scale up the consumer to 2
consumer-1 consumes data at 500 ms intervals
consumer-2 consumes data at 300 ms intervals
Print the data with the consumer id
*/

// consumer
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	go genNos(ch)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("consumer-1 :", val)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("consumer-2 :", val)
			time.Sleep(300 * time.Millisecond)
		}
	}()
	wg.Wait()
}

// producer
func genNos(ch chan int) {
	count := rand.Intn(20)
	fmt.Printf("[genNos] producing %d values\n", count)
	for idx := 1; idx <= count; idx++ {
		ch <- 10 * idx
	}
	close(ch)
}
