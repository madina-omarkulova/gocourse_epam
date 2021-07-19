package main

import (
	"fmt"

	"github.com/madina-omarkulova/gocourse_epam/lesson04/shapes"
)

func DescribeShape(s shapes.Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// choose your own dimensions
	c := shapes.Circle{
		radius: 5,
	}
	r := shapes.Rectangle{
		height: 10,
		width:  25,
	}
	DescribeShape(c)
	DescribeShape(r)
}
