package main

import (
	"fmt"
	"sync"
)

/*
func main(){
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
*/

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(x int) {
			fmt.Println(x)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
