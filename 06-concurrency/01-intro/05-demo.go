package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	var count, cores int

	flag.IntVar(&count, "count", 0, "# of goroutines to run")
	flag.IntVar(&cores, "cores", runtime.GOMAXPROCS(0), "# of cpu cores to use")
	flag.Parse()
	runtime.GOMAXPROCS(cores)  //set
	x := runtime.GOMAXPROCS(0) //get (without changing the prev value)
	fmt.Printf("Starting %d goroutines using %d cores.. hit ENTER to start!\n", count, x)

	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1)    // increment the wg counter by 1
		go fn(i, wg) //=> schedule the execution of fn through the scheduler
	}
	wg.Wait() //=> block until the counter becomes 0 (default)
	fmt.Println("Done!!")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
