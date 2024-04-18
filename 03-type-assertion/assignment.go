package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10))                      //=> 10
	fmt.Println(sum(10, 20))                  //=> 30
	fmt.Println(sum(10, "20", 30, 40))        //=> 100
	fmt.Println(sum(10, "20", 30, 40, "abc")) //=> 100
	fmt.Println(sum())                        //=> 0
}

func sum(values ...interface{}) int {
	var result int
	for _, val := range values {
		switch n := val.(type) {
		case int:
			result += n
		case string:
			if x, err := strconv.Atoi(n); err == nil {
				result += x
			}
		}
	}
	return result
}
