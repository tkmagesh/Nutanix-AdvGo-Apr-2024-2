package main

import (
	"fmt"
	"math"
)

// sprint-1
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// sprint-2
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// sprint-3
/*
func PrintArea(x interface{}) {
	switch obj := x.(type) {
	case Circle:
		fmt.Println("Area :", obj.Area())
	case Rectangle:
		fmt.Println("Area :", obj.Area())
	default:
		fmt.Println("Cannot calculate area!")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch obj := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", obj.Area())
	default:
		fmt.Println("type does not have Area() method")
	}
}
*/

type AreaFinder interface{ Area() float64 }

func PrintArea(obj AreaFinder) {
	fmt.Println("Area :", obj.Area())
}

// sprint-4
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type PerimeterFinder interface{ Perimeter() float64 }

func PrintPerimeter(obj PerimeterFinder) {
	fmt.Println("Perimeter :", obj.Perimeter())
}

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

// sprint-5
func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)      //=> interface{Area() float64}
	PrintPerimeter(x) //=> interface{Perimeter() float64}
}

func main() {
	c := Circle{Radius: 10}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShapeStats(c)

	r := Rectangle{Height: 12, Width: 10}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShapeStats(r)

}
