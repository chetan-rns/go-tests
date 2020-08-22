package structs

import (
	"math"
	"testing"
)

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.height + r.width)
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{width: 2.0, height: 5.0})

	want := 14.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func Area(r Rectangle) float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	height float64
	base   float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func TestArea(t *testing.T) {

	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle",
			shape: &Rectangle{2.0, 5.0},
			want:  10,
		},
		{
			name:  "Circle",
			shape: &Circle{2.0},
			want:  12.566370614359172,
		},
		{
			name:  "Triangle",
			shape: &Triangle{12, 6},
			want:  36,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("%#v got %g want %g", test.shape, got, test.want)
			}
		})
	}
}
