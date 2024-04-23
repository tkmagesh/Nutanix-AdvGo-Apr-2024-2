package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data1.dat")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(txt)
	}
}
