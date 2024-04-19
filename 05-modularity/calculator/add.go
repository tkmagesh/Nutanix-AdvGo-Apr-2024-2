package calculator

import "fmt"

func init() {
	fmt.Println("calculator.init() [add.go] invoked")
}

func Add(x, y int) int {
	opCount["Add"]++
	return x + y
}
