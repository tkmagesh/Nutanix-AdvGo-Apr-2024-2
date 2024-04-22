package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	dataCh := genData(rootCtx)
	for data := range dataCh {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genData(ctx context.Context) <-chan string {
	wg := &sync.WaitGroup{}
	dataCh := make(chan string)

	go func() {
		cancelCtx, cancel := context.WithCancel(ctx)
		fmt.Println("Hit ENTER to stop generating fib series....")

		go func(cancel context.CancelFunc) {
			fmt.Scanln()
			cancel()
		}(cancel)

		timeoutCtx, cancel := context.WithTimeout(cancelCtx, 5*time.Second)
		defer cancel()

		fibCh := genFib(cancelCtx)
		primeCh := genPrime(timeoutCtx)

		wg.Add(1)
		go printFib(wg, fibCh, dataCh)

		wg.Add(1)
		go printPrime(wg, primeCh, dataCh)

		wg.Wait()
		close(dataCh)
	}()

	return dataCh

}

func printFib(wg *sync.WaitGroup, fibCh <-chan int, dataCh chan<- string) {
	defer wg.Done()
	for fibNo := range fibCh {
		dataCh <- fmt.Sprintf("fib : %d", fibNo)
	}
}

func printPrime(wg *sync.WaitGroup, PrimeCh <-chan int, dataCh chan<- string) {
	defer wg.Done()
	for primeNo := range PrimeCh {
		dataCh <- fmt.Sprintf("prime : %d", primeNo)
	}
}

// using "share memory by communicating" (advisable)
func genPrime(ctx context.Context) <-chan int {
	primeCh := make(chan int)
	go func() {
	LOOP:
		for no := 2; ; no++ {
			select {
			case <-ctx.Done():
				fmt.Println("stop signal received... closing the prime channel")
				close(primeCh)
				break LOOP
			default:
				if !isPrime(no) {
					continue LOOP
				}
				primeCh <- no
				time.Sleep(500 * time.Millisecond)
			}
		}

	}()
	return primeCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func genFib(ctx context.Context) <-chan int {
	fibCh := make(chan int)
	go func() {
	LOOP:
		for x, y := 0, 1; ; x, y = y, x+y {
			select {
			case <-ctx.Done():
				fmt.Println("stop signal received... closing the fib channel")
				close(fibCh)
				break LOOP
			case fibCh <- x:
				time.Sleep(500 * time.Millisecond)
			}
		}

	}()
	return fibCh
}
