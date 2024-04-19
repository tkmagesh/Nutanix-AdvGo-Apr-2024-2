package main

import (
	"fmt"

	"github.com/fatih/color"
	calc "github.com/tkmagesh/nutanix-advgo-apr-2024-2/05-modularity/calculator"
)

func main() {
	color.Red("Hello, World!")
	greet("Magesh")
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Op Count :", calc.OpCount())
}
