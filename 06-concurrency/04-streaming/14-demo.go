package main

import (
	"fmt"
	"math/rand"
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
	go genNos(ch)
	for val := range ch {
		fmt.Println(val)
		time.Sleep(500 * time.Millisecond)
	}
}

// producer
func genNos(ch chan int) {
	count := rand.Intn(100)
	fmt.Printf("[genNos] producing %d values\n", count)
	for idx := 1; idx <= count; idx++ {
		ch <- 10 * idx
	}
	close(ch)
}
