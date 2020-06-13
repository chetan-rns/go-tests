package structs

import (
	"math"
	"testing"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func TestArea(t *testing.T) {

	tt := []struct {
		name string
		shape
		want float64
	}{
		{"area of a rectangle", rectangle{2.0, 2.0}, 4.0},
		{"area of a circle", circle{10.0}, 314.1592653589793},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.area()
			checkValue(t, got, test.want)
		})
	}
}

func checkValue(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
