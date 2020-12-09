package shapes

import "testing"

func TestPerimeter(t *testing.T) {
    rectangle := Rectange{10.0, 10.0}
    got := rectangle.Perimeter()
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {

    tests := []struct {
        name string
        shape Shape
        want float64
    }{
        {name: "Rectangle", shape: Rectange{Width: 12, Height: 6}, want: 72.0},
        {name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Width: 12, Height: 6}, want: 36.0},
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
