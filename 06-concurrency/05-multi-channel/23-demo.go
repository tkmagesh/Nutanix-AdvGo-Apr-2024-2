/* multiple producers & one consumer */

/*
producer-1 : generates 20 values in fib series at 500 ms intervals
producer-2 : generates 40 odd numbers at 300 ms intervals
consume them and print them
*/

/*
refactor the below using select-case
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	wg := &sync.WaitGroup{}

	go func(ch chan string) {
		fibCh := make(chan int)
		wg.Add(1)
		go genFib(fibCh, wg)
		defer wg.Done()
		for fibNo := range fibCh {
			ch <- fmt.Sprintf("fib : %d", fibNo)
		}
	}(ch)

	go func(ch chan string) {
		oddCh := make(chan int)
		wg.Add(1)
		go genOdd(oddCh, wg)
		defer wg.Done()
		for oddNo := range oddCh {
			ch <- fmt.Sprintf("oddNo : %d", oddNo)
		}
	}(ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}

}

func genFib(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x, y, count := 0, 1, 1; count <= 20; count, x, y = count+1, y, x+y {
		time.Sleep(500 * time.Millisecond)
		ch <- x
	}

}

func genOdd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for no, count := 1, 1; count <= 40; count, no = count+1, no+2 {
		ch <- no
		time.Sleep(300 * time.Millisecond)
	}

}
