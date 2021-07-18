package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c Circle) Area() float64 {
	return (math.Pi * c.radius * c.radius)
}

func (r Rectangle) Area() float64 {
	return (r.height * r.width)
}

func (c Circle) Perimeter() float64 {
	return (math.Pi * c.radius * 2)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
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

func (s Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", s.radius)
}

func (s Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f ", s.height, s.width)
}
