package shapes

import (
	"math"
)

// Shape is anything that has Area and Perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle represents a circle
type Circle struct {
	Radius float64
}

// Rectangle represents a rectangle
type Rectangle struct {
	Height float64
	Width  float64
}

// Triangle represents a triangle
type Triangle struct {
	Base   float64
	Height float64
	Slope  float64
}

// Area returns the area of the rectangle
func (r Rectangle) Area() (result float64) {
	result = r.Height * r.Width
	return
}

// Perimeter returns the perimeter of the rectangle
func (r Rectangle) Perimeter() (result float64) {
	result = (r.Height + r.Width) * 2
	return
}

// Area returns the area of the circle
func (c Circle) Area() (result float64) {
	result = math.Pi * c.Radius * c.Radius
	return
}

// Perimeter returns the perimeter of the circle
func (c Circle) Perimeter() (result float64) {
	result = math.Pi * c.Radius * 2
	return
}

// Area returns the area of the triangle
func (t Triangle) Area() (result float64) {
	result = t.Base * t.Height * 0.5
	return
}

// Perimeter returns the perimeter of the triangle
func (t Triangle) Perimeter() (result float64) {
	result = t.Base + t.Height + t.Slope
	return
}
