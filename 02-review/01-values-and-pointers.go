// Everything is a value in go

package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	// var p1 Product
	// fmt.Println(p1)
	p1 := Product{Id: 100, Name: "Pen", Cost: 5}
	p2 := p1 // a copy is created because everything is a value
	p2.Cost = 100
	fmt.Printf("p1 : %#v\n", p1)
	fmt.Printf("p2 : %#v\n", p2)

	// comparison by value
	apsaraPencil := Product{Id: 200, Name: "Pencil", Cost: 10}
	natrajPencil := Product{Id: 200, Name: "Pencil", Cost: 10}
	fmt.Println(apsaraPencil == natrajPencil)

	fmt.Printf("&apsaraPencil : %p\n", &apsaraPencil)
	fmt.Printf("&natrajPencil : %p\n", &natrajPencil)

	// pointers
	var no int = 100
	// noPtr := &no // value -> reference
	var noPtr *int
	noPtr = &no
	fmt.Println(noPtr)
	// dereference (reference -> value)
	var x int
	x = *noPtr
	fmt.Println(x)

	// struct pointers
	prod := Product{Id: 100, Name: "Pen", Cost: 5}
	prodPtr := &prod
	fmt.Println("prod.Id =", prod.Id)
	fmt.Println("prodPtr.Id =", prodPtr.Id)

	// Using struct "value"
	/*
		fmt.Printf("[main] &prod = %p\n", &prod)
		fmt.Println(prod.Format())
		fmt.Println("After applying 10% discount")
		prod.ApplyDiscount(10)
		fmt.Println(prod.Format())
	*/
	fmt.Println()
	fmt.Println("Using struct 'pointer'")
	fmt.Printf("[main] &prodPtr = %p\n", prodPtr)
	fmt.Println(prodPtr.Format())
	fmt.Println("After applying 10% discount")
	prodPtr.ApplyDiscount(10)
	fmt.Println(prodPtr.Format())

}

func (p Product) Format() string {
	fmt.Printf("[Format] &prod = %p\n", &p)
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercent float32) {
	fmt.Printf("[ApplyDiscount] &prod = %p\n", p)
	p.Cost = p.Cost * ((100 - discountPercent) / 100)
}
