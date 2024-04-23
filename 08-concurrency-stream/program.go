package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	fileWg := &sync.WaitGroup{}
	dataCh := make(chan int)

	fileWg.Add(1)
	go Source(fileWg, "data1.dat", dataCh)

	fileWg.Add(1)
	go Source(fileWg, "data2.dat", dataCh)

	oddCh, evenCh := Splitter(dataCh)
	oddSumCh := Sum(oddCh)
	evenSumCh := Sum(evenCh)

	done := Merger(oddSumCh, evenSumCh)

	fileWg.Wait()
	close(dataCh)

	<-done
	fmt.Println("Done")
}

func Source(wg *sync.WaitGroup, fileName string, ch chan<- int) {
	defer wg.Done()
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if val, err := strconv.Atoi(txt); err == nil {
			ch <- val
		}
	}
}

func Splitter(dataCh <-chan int) (<-chan int, <-chan int) {
	oddCh := make(chan int)
	evenCh := make(chan int)
	go func() {
		defer close(oddCh)
		defer close(evenCh)
		for data := range dataCh {
			if data%2 == 0 {
				evenCh <- data
			} else {
				oddCh <- data
			}
		}
	}()
	return oddCh, evenCh
}

func Sum(ch <-chan int) <-chan int {
	sumCh := make(chan int)
	go func() {
		defer close(sumCh)
		var result int
		for val := range ch {
			result += val
		}
		fmt.Println("result :", result)
		sumCh <- result
	}()
	return sumCh
}

func Merger(oddSumCh, evenSumCh <-chan int) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		file, err := os.OpenFile("result.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		for i := 0; i < 2; i++ {
			select {
			case evenSum := <-evenSumCh:
				fmt.Printf("Even Total : %d\n", evenSum)
				fmt.Fprintf(file, "Even Total : %d\n", evenSum)
			case oddSum := <-oddSumCh:
				fmt.Printf("Odd Total : %d\n", oddSum)
				fmt.Fprintf(file, "Odd Total : %d\n", oddSum)
			}
		}
		close(doneCh)
	}()
	return doneCh
}
