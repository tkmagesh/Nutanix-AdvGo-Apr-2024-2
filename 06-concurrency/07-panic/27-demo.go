package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic occurred :", err)
				return
			}
			fmt.Println("Thank you!")
		}()
	*/
	/*
		ch, errCh := divide(100, 0)
		select {
		case result := <-ch:
			fmt.Println(result)
		case err := <-errCh:
			fmt.Println("error occurred :", err)
		}
	*/

	ch, _ := divide(100, 0)
	fmt.Println(<-ch)

}

func divide(x, y int) (<-chan int, <-chan error) {
	ch := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				errCh <- err.(error)
			}
		}()
		time.Sleep(2 * time.Second)
		ch <- x / y
	}()
	return ch, errCh
}
