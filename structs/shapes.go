package main

import (
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Triangle struct {
	Base   float64
	Height float64
	Slope  float64
}

func (r Rectangle) Area() (result float64) {
	result = r.Height * r.Width
	return
}

func (r Rectangle) Perimeter() (result float64) {
	result = (r.Height + r.Width) * 2
	return
}

func (c Circle) Area() (result float64) {
	result = math.Pi * c.Radius * c.Radius
	return
}

func (c Circle) Perimeter() (result float64) {
	result = math.Pi * c.Radius * 2
	return
}

func (t Triangle) Area() (result float64) {
	result = t.Base * t.Height * 0.5
	return
}

func (t Triangle) Perimeter() (result float64) {
	result = t.Base + t.Height + t.Slope
	return
}
