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

	fmt.Println("Adding a new key")
	productRanks["mouse"] = 2
	fmt.Println(productRanks)

	fmt.Println("Iterating using range")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Check for the existence of a key")
	if rank, exists := productRanks["scribble-pad"]; exists {
		fmt.Println("Rank of scribble-pad :", rank)
	} else {
		fmt.Println("scribble-pad is not ranked yet!")
	}

	fmt.Println("Removing an item")
	keyToRemove := "scribble-pad"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)

}
