package main

import "fmt"

var counter int

func main() {
	for i := 1; i <= 200; i++ {
		// execute the increment() function concurrently
		increment()
	}
	fmt.Println(counter)
}

func increment() {
	counter++
}
