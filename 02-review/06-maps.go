/* maps -> typed collection of key/value pairs */
package main

import "fmt"

func main() {
	// var productRanks map[string]int = make(map[string]int)
	// var productRanks map[string]int = map[string]int{}
	// productRanks := map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	productRanks := map[string]int{
		"pen":    5,
		"pencil": 1,
		"marker": 3,
	}

	fmt.Println("len(productRanks) :", len(productRanks))

	fmt.Println("Iterating using range")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}
}
