package shapes

import (
	"testing"
)

func TestShape(t *testing.T) {

	shapeTests := []struct {
		name         string
		shape        Shape
		hasArea      float64
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0, hasPerimeter: 36},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{3, 4, 5}, hasArea: 6.0, hasPerimeter: 12},
	}

	check := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %g expected %g", got, want)
		}
	}

	for _, tt := range shapeTests {
		t.Run(tt.name+" area", func(t *testing.T) {
			got := tt.shape.Area()
			check(t, got, tt.hasArea)
		})

		t.Run(tt.name+" perimeter", func(t *testing.T) {
			got := tt.shape.Perimeter()
			check(t, got, tt.hasPerimeter)
		})
	}
}
