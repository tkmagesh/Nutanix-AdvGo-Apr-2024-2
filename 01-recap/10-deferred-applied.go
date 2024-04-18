package main

import "fmt"

func main() {
	q, r, err := divideWrapper(100, 0)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	fmt.Println(q, r)
}

func divideWrapper(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api
func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
