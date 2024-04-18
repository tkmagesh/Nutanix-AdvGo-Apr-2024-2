package main

import "fmt"

func main() {
	counter := getCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func getCounter() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}
