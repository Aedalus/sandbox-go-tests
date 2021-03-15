package main

import "testing"

func TestPerimeter(t *testing.T){
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want float64
	} {
		{Rectangle{Width: 12,Height: 6}, 72.0},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Width: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}