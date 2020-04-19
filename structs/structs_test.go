package structs

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{2.0, 3.0}
	got := rectangle.Permeter()
	want := 10.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{2.0, 3.0}
	got := rectangle.Area()
	want := 6.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Permeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}
