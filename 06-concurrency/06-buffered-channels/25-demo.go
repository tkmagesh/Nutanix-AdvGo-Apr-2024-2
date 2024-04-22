package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n ", len(ch), cap(ch))
	ch <- 100
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n ", len(ch), cap(ch))
	ch <- 200
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n ", len(ch), cap(ch))
	fmt.Println("data from ch :", <-ch)
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n ", len(ch), cap(ch))
	fmt.Println("data from ch :", <-ch)
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n ", len(ch), cap(ch))
}
