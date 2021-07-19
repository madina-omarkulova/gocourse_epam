package Shapes

import "fmt"

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return (r.height * r.width)
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (s Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f ", s.height, s.width)
}
