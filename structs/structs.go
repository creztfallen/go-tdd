package structs

import "math"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

type Shape interface {
	Area() float64
}
