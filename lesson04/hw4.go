package main

import (
	"fmt"
	"math"
)

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

func (s Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", s.radius)
}

func (s Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f ", s.height, s.width)
}
