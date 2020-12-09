package shapes

import "math"

type Shape interface {
    Area() float64
}

type Rectange struct {
    Width float64
    Height float64
}

func (r Rectange) Area() float64 {
    return r.Width * r.Height
}

func (r Rectange) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
    Width float64
    Height float64
}

func (t Triangle) Area() float64 {
    return (t.Width * t.Height) / 2
}

