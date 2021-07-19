package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}
type Rectangle struct {
	height float64
	width  float64
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// choose your own dimensions
	c := Circle{radius: 5}
	r := Rectangle{
		height: 10,
		width:  25,
	}
	DescribeShape(c)
	DescribeShape(r)
}
