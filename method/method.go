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
	return a.x * a.y * 2
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
	areaKvadrat := square{4}
	fmt.Printf("Kvadrat tomoni=%0.0f \nKvadrat Yuzasi=", areaKvadrat)
	fmt.Println(areaKvadrat.SquareArea())
	fmt.Printf("Kvadrat Perimetri=%0.0f", areaKvadrat)
	fmt.Println(areaKvadrat.SquarePer())
}
