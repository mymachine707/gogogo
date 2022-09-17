package main

import "fmt"

type Cube struct {
	a float32
}

// Perimeter ...
func (obj Cube) Perimeter() float64 {
	return float64(12 * obj.a)
}

// Area ...
func (obj Cube) Area() float64 {
	return float64(6 * obj.a * obj.a)
}

// Area ...
func (obj Cube) Volume() float64 {
	return float64(obj.a * obj.a * obj.a)
}

// Shape...
type Shape interface {
	Perimeter() float64
	Area() float64
}

// Object ...
type Object interface {
	Volume() float64
}

func main() {
	cube := Cube{
		a: 4,
	}

	var shape Shape
	shape = cube

	fmt.Printf("Area = %f, Perimeter = %f\n", shape.Area(), shape.Perimeter())

	var obj Object
	obj = cube

	fmt.Printf("Volume = %f\n", obj.Volume())
}
