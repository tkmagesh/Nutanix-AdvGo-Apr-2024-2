package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		if val, isOpen := <-ch; isOpen {
			fmt.Println(val)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}

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
