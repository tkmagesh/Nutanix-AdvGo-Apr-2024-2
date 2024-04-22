package main

import (
	"fmt"
)

func main() {

	defer func() {
		fmt.Println("main[deferred]")
		if err := recover(); err != nil {
			fmt.Println("panic occurred :", err)
			return
		}
		fmt.Println("Thank you!")
	}()

	ch, errCh := divide(100, 0)
	select {
	case result := <-ch:
		fmt.Println(result)
	case err := <-errCh:
		fmt.Println("error occurred :", err)
	}

	/*
		ch, _ := divide(100, 0)
		fmt.Println(<-ch)
	*/

}

func divide(x, y int) (<-chan int, <-chan error) {
	resultCh := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		defer func() {
			fmt.Println("divide[deferred]")
			if err := recover(); err != nil {
				errCh <- err.(error)
			}
		}()
		ch := divideApi(x, y)
		resultCh <- <-ch
	}()
	return resultCh, errCh

}

func divideApi(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		ch <- x / y
	}()
	return ch
}
