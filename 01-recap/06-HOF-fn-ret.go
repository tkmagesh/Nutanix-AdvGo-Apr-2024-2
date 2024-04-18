package main

import "fmt"

func main() {
	fn := getFn()
	fn()
}

func getFn() func() {
	/*
		var fn func()
		fn = func() {
			fmt.Println("fn invoked")
		}
		return fn
	*/
	return func() {
		fmt.Println("fn invoked")
	}
}
