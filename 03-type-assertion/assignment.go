package main

import "fmt"

func main() {
	fmt.Println(sum(10)) //=> 10
	fmt.Println(sum(10, 20)) //=> 30
	fmt.Println(sum(10, "20", 30, 40)) //=> 100
	fmt.Println(sum(10, "20", 30, 40, "abc")) //=> 100
	fmt.Println(sum()) //=> 0
}

func sum(?) int {
	var result int
	?
	return result
}
