package main

import "fmt"

func main() {
	// var nos []int
	var nos = make([]int, 0, 4)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 10)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 20)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 30)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 40)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 50)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

}
