package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	timeOutCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	fmt.Println("Hit ENTER to stop or will shutdown after 5 secs")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	ch := genFib(timeOutCtx)
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println(timeOutCtx.Err().Error())
	fmt.Println("Done")
}

func genFib(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for x, y := 0, 1; ; x, y = y, x+y {
			select {
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
			case <-ctx.Done():
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}

/*
func genOdd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for no, count := 1, 1; count <= 40; count, no = count+1, no+2 {
		ch <- no
		time.Sleep(300 * time.Millisecond)
	}

} */
