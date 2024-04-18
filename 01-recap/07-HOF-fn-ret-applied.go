package main

import "fmt"

func main() {
	add := getLogOperation(add)
	add(100, 200)

	subtract := getLogOperation(subtract)
	subtract(100, 200)

	/*
		logAdd(100,200)
		logSubtract(100,200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		logOperation(func(x, y int) {
			fmt.Println("Multiply result :", x*y)
		}, 100, 200)
	*/

	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)

		logMultiply := getLogOperation(func(x, y int) {
			fmt.Println("Multiply result :", x*y)
		})
		logMultiply(100, 200)
	*/
}

func getLogOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("Operation started")
		op(x, y)
		fmt.Println("Operation completed")
	}
}

/*
func logOperation(op func(int, int), x, y int) {
	fmt.Println("Operation started")
	op(x, y)
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
