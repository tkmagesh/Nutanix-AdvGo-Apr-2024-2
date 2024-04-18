package main

import "fmt"

func main(){
	/* 
	exec("f1")
	exec("f2")
	exec("F1") 
	*/

	exec(f1)
	exec(f2)
	// exec(F1)
	exec(func(){
		fmt.Println("anon fn invoked")
	})
	exec(func(){
		add(100,200)
	})
}

func add(x,y int){
	fmt.Println(x + y)
}

/* 
func exec(fnName string){
	// execute either f1 or f2 based on the argument
	switch fnName {
	case "f1":
		f1()
	case "f2":
		f2()
	default:
		panic("invalid argument")
	}
} 
*/

func exec(fn func()){
	fn()
}



func f1(){
	fmt.Println("f1 invoked")
}

func f2(){
	fmt.Println("f2 invoked")
}
