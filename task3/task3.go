package main

import (
	"fmt"
)

func main() {
	var r float32 = 10.04

	// fmt.Scanf("%f", &r)
	fmt.Println("R =", r)

	fmt.Printf("Area: %0.2f\n", area(r))
}

func area(r float32) (area float32) {
	//
	// WRITE YOUR CODE HERE
	const p float32 = 3.14
	var circle float32 = p * r * r
	var square float32 = (r + r) * (r + r)
	area = square - circle

	//
	return area
}
