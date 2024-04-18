/*
Array - Fixed sized typed collection
*/
package main

import "fmt"

func main() {
	// var nos [5]int // memory allocated & initialized
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	var nos [5]int = [5]int{3, 1, 4}
	fmt.Println(nos)

	// iteration using indexer
	fmt.Println("iteration using indexer")
	for idx := 0; idx < 5; idx++ {
		fmt.Printf("nos[%d] : %d\n", idx, nos[idx])
	}

	fmt.Println("iteration using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] : %d\n", idx, val)
	}

	fmt.Println("assignment behavior")
	x := [3]int{10, 20, 30}
	y := [3]int{10, 20, 30}
	fmt.Println(x == y)

	nos2 := nos // creates a copy coz "everything is a value"
	nos2[4] = 9999
	fmt.Println(nos[4], nos2[4])

	fmt.Println("After doubling the list")
	doubleList(nos)
	for idx, val := range nos {
		fmt.Printf("nos[%d] : %d\n", idx, val)
	}
}

func doubleList(list [5]int) {
	for idx, val := range list {
		list[idx] = val * 2
	}
}
