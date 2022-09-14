package main

import (
	"fmt"
	"math"
)

// -----------------------------------------// -----------------------------------------// -----------------------------------------

// Square//////////////// Kvadrat

type square struct {
	x float64
}

func (a square) SquareArea() float64 {
	// x tomon uzunligi
	return a.x * a.x
}

func (a square) SquarePer() float64 {
	return 4 * a.x
}

// Rectangle///////////// To'rtburchak// -----------------------------------------// -----------------------------------------// -----------------------------------------

type Rectangle struct {
	x, y float64
}

func (a Rectangle) RectangleArea() float64 {
	// x tomon uzunligi
	return a.x * a.y
}

func (a Rectangle) RectanglePer() float64 {
	return (a.x + a.y) * 2
}

// Circle ///////////////// aylana// -----------------------------------------// -----------------------------------------// -----------------------------------------

type Circle struct {
	x float64
}

func (a Circle) CircleArea() float64 {
	// x yuza uzunligi
	return a.x * a.x * math.Pi
}

func (a Circle) CirclePer() float64 {
	return 2 * a.x * math.Pi
}

// Triangle /////////////-------------------// -----------------------------------------// -----------------------------------------// -----------------------------------------

// Triangle

type Triangle struct {
	x, y, z float64
}

func (a Triangle) TriangleArea() float64 {
	// x yuza uzunligi
	p := (a.x + a.y + a.z) / 2
	S := math.Sqrt(p * (p - a.x) * (p - a.y) * (p - a.z))
	return S
}

func (a Triangle) TrianglePer() float64 {
	p := (a.x + a.y + a.z) / 2
	return p
}

// Triangle /////////////-------------------// -----------------------------------------// -----------------------------------------// -----------------------------------------

func main() {
	areaKvadrat := square{4} // kvadrat
	fmt.Printf("Kvadrat tomoni=%0.0f \nKvadrat Yuzasi=", areaKvadrat)
	fmt.Println(areaKvadrat.SquareArea())
	fmt.Printf("%s", "Kvadrat Perimetri=")
	fmt.Println(areaKvadrat.SquarePer())

	fmt.Println("\n")

	Rectangles := Rectangle{4, 6} // To'rt burchak
	fmt.Printf("To'rt burchak tomonlarii=%0.0f \nTo'rt burchak Yuzasi=", Rectangles)
	fmt.Println(Rectangles.RectangleArea())
	fmt.Printf("%s", "To'rt burchak Perimetri=")
	fmt.Println(Rectangles.RectanglePer())

	fmt.Println("\n")

	Circles := Circle{4} // Aylana
	fmt.Printf("Aylana radiusi=%0.0f \nAylana Yuzasi=", Circles)
	fmt.Println(Circles.CircleArea())
	fmt.Printf("%s", "Aylana yoyi uzunligi=")
	fmt.Println(Circles.CirclePer())

	fmt.Println("\n")

	Triangles := Triangle{3, 4, 5} // Uch burchak
	fmt.Printf("Uch burchak tomonlarii=%0.0f \nUch burchak Yuzasi=", Triangles)
	fmt.Println(Triangles.TriangleArea())
	fmt.Printf("%s", "Uch burchak Perimetri=")
	fmt.Println(Triangles.TrianglePer())

	fmt.Println("\n")
}
