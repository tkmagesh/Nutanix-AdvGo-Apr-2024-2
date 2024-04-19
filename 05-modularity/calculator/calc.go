package calculator

import "fmt"

var opCount map[string]int

func init() {
	fmt.Println("calculator.init() [calc.go] invoked")
	opCount = make(map[string]int)
}

func OpCount() map[string]int {
	return opCount
}
