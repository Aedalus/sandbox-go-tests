package main

import "math"

type Shape interface {
	Area() float64
}

// Rectangle
type Rectangle struct {
	Width float64
	Height float64
}
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Triangle
type Triangle struct {
	Width float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Width * t.Height * 0.5
}

