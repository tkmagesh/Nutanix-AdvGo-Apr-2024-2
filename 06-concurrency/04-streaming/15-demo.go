/* receive operation on a closed channel */
package main

func main() {
	ch := make(chan int)
	close(ch)
	ch <- 100
}
