package main

import (
	"fmt"

	"github.com/madina-omarkulova/gocourse_epam/lesson04/Shapes"
)

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// choose your own dimensions
	c := Shapes.Circle{
		radius: 5,
	}
	r := Shapes.Rectangle{
		height: 10,
		width:  25,
	}
	DescribeShape(c)
	DescribeShape(r)
}
