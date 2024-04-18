package main

import "fmt"

func main(){
	var fn func()
	fn = func (){
		fmt.Println("fn invoked - 1")
	}
	fn()

	fn = func (){
		fmt.Println("fn invoked - 2")
	}
	fn()

	var greeter func(string)
	greeter = func(userName string){
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greeter("Magesh")
	
	var add func(int, int)
	add = func(x,y int){
		fmt.Println(x + y)
	}
	add(200,200)

	var multiply func(int, int) int
	multiply = func(x,y int) int {
		return x * y
	}
	result := multiply(100,200)
	fmt.Println(result)

	var divide func(int, int)(int, int)
	divide = func(x,y int) (quotient, remainder int){
		quotient, remainder = x/y, x % y
		return
	}
	q, r := divide(100,7)
	fmt.Println(q, r) 
	
}