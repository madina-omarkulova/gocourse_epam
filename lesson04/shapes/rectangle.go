package shapes

import "fmt"

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (s Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f ", s.Height, s.Width)
}
