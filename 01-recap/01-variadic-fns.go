package main

import "fmt"

func main(){
	fmt.Println(sum(10))
	fmt.Println(sum(10,20))
	fmt.Println(sum(10,20,30,40))
	fmt.Println(sum())
}

func sum(values ...int) int {
	var result int
	for _, val := range values {
		result += val
	}
	return result
}