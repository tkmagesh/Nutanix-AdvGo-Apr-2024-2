/* multiple producers & one consumer */

/*
producer-1 : generates 20 values in fib series at 500 ms intervals
producer-2 : generates 40 odd numbers at 300 ms intervals
consume them and print them
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go genFib(ch, wg)

	wg.Add(1)
	go genOdd(ch, wg)

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
