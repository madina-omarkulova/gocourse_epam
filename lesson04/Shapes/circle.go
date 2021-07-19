package Shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return (math.Pi * c.radius * c.radius)
}

func (c Circle) Perimeter() float64 {
	return (math.Pi * c.radius * 2)
}
func (s Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", s.radius)
}
