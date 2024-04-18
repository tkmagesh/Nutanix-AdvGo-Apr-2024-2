package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100,200)
		logSubtract(100,200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(x, y int) {
		fmt.Println("Multiply result :", x*y)
	}, 100, 200)
}

func logOperation(op func(int, int), x, y int) {
	fmt.Println("Operation started")
	op(x, y)
	fmt.Println("Operation completed")
}

/*
func logAdd(x, y int) {
	fmt.Println("Operation started")
	add(x, y)
	fmt.Println("Operation completed")
}

func logSubtract(x, y int) {
	fmt.Println("Operation started")
	subtract(x, y)
	fmt.Println("Operation completed")
}
*/
// 3rd party library api
func add(x, y int) {
	fmt.Println("Add result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result :", x-y)
}
