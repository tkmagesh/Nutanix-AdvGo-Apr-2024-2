package main

import "fmt"

func main() {
	ch := make(chan int)
	/*
		data := <-ch
		ch <- 100
		fmt.Println(data)
	*/

	/*
		 	ch <- 100
			data := <-ch
			fmt.Println(data)
	*/
	/*
		go func() {
			ch <- 100
		}()
		data := <-ch
		fmt.Println(data)
	*/
	go func() {
		data := <-ch
		fmt.Println(data)
	}()
	ch <- 100

}
