package shapes

import (
	"math"
)

// Shape interface representing a 2D shape
type Shape interface {
	Area() float64
	Perim() float64
}

// Rectangle object
type Rectangle struct {
	Height float64
	Width  float64
}

// Area of a given rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Perim computes the perimeter of a rectangle
func (r Rectangle) Perim() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle object
type Circle struct {
	Radius float64
}

// Area of a given circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Perim computes the perimeter of a circle
func (c Circle) Perim() float64 {
	return math.Pi * 2 * c.Radius
}
