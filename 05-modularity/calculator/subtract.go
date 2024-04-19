package calculator

import "fmt"

func init() {
	fmt.Println("calculator.init() [subtract.go] invoked")
}

func Subtract(x, y int) int {
	opCount["Subtract"]++
	return x - y
}
