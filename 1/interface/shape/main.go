// Circle, Triangle, Square, Rect
// () Perimeter() flaot64
// () Area() float64

package main

import (
	"errors"
	"fmt"
	"math"
)

// Circle ...
type Circle struct {
	r float32
}

// Perimeter ...
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * float64(c.r)
}

// Area ...
func (c Circle) Area() float64 {
	return math.Pi * float64(c.r*c.r)
}

// Triangle ...
type Triangle struct {
	a, b, c float32
}

// Perimeter ...
func (t Triangle) Perimeter() float64 {
	return float64(t.a + t.b + t.c)
}

// Area ...
func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2
	area := math.Sqrt(s * (s - float64(t.a)) * (s - float64(t.b)) * (s - float64(t.c)))
	return area
}

// Check ...
func (t Triangle) Check() error {
	var err error
	if t.a+t.b <= t.c || t.a+t.c <= t.b || t.c+t.b <= t.a {
		err = errors.New("Triangle is not valid")
		return err
	}

	return err
}

// Square ...
type Square struct {
	a float32
}

// Perimeter ...
func (obj Square) Perimeter() float64 {
	return float64(4 * obj.a)
}

// Area ...
func (obj Square) Area() float64 {
	return float64(obj.a * obj.a)
}

// Rect ...
type Rect struct {
	a, b float32
}

// Perimeter ...
func (obj Rect) Perimeter() float64 {
	return float64(2 * (obj.a + obj.b))
}

// Area ...
func (obj Rect) Area() float64 {
	return float64(obj.a * obj.b)
}

func main() {
	obj1 := Triangle{
		a: 3,
		b: 4,
		c: 5,
	}

	Calculate(obj1)

	obj2 := Circle{
		r: 3,
	}

	Calculate(obj2)

	obj3 := Square{
		a: 3,
	}
	Calculate(obj3)

	obj4 := Rect{
		a: 3,
		b: 4,
	}
	Calculate(obj4)
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

// Calculate ...
func Calculate(obj Shape) {
	switch obj.(type) {
	case Circle:
		fmt.Println("Circle =>")
	case Triangle:
		fmt.Println("Triangle =>")
	case Square:
		fmt.Println("Square =>")
	default:
		fmt.Println("Unknown =>")
	}

	fmt.Printf("Area = %f, Perimeter = %f\n", obj.Area(), obj.Perimeter())
	fmt.Println()
}
