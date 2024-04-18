// Slice - Varying sized typed collection
package main

import "fmt"

func main() {
	// var nos []int
	var nos []int = []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// appending new values
	nos = append(nos, 10)
	nos = append(nos, 10, 20, 30, 40)

	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)

	// iteration using indexer
	fmt.Println("iteration using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] : %d\n", idx, nos[idx])
	}

	fmt.Println("iteration using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] : %d\n", idx, val)
	}

	nos2 := nos
	nos2[0] = 9999
	fmt.Println(nos[0], nos2[0])

	fmt.Println("After doubling the list")
	doubleList(nos)
	for idx, val := range nos {
		fmt.Printf("nos[%d] : %d\n", idx, val)
	}

}

func doubleList(list []int) {
	for idx, val := range list {
		list[idx] = val * 2
	}
}
