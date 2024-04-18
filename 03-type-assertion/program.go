package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = "Irure quis exercitation qui irure pariatur eu amet proident."
	x = true
	x = 99.98
	x = 5 + 7i
	fmt.Println(x)

	x = 200
	// x = "Et nulla reprehenderit ad commodo est eiusmod consequat voluptate ut Lorem commodo proident ea."
	// y := x * 100
	y := x.(int) * 200 // can lead to a panic
	fmt.Println(y)

	// runtime safety
	x = "Sint voluptate laboris reprehenderit cupidatat duis."
	if val, ok := x.(int); ok {
		result := val * 2
		fmt.Println(result)
	} else {
		fmt.Println("x is not an int, so can't multiply")
	}

	// using type-switch
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case float64:
		fmt.Println("x is a float64, x * .99 =", val*0.99)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	default:
		fmt.Println("Unknown type")
	}
}
