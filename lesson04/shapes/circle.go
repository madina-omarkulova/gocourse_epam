package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}

func (c Circle) Perimeter() float64 {
	return (math.Pi * c.Radius * 2)
}
func (s Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", s.Radius)
}
