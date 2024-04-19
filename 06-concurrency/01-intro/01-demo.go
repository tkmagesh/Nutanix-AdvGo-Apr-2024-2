package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //=> schedule the execution of f1 through the scheduler
	f2()

	// DO NOT use time.Sleep() for goroutine synchronization
	time.Sleep(2 * time.Second)
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
