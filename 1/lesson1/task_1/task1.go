package main

import "fmt"

func main() {
	var a, b int = 3, 4

	// fmt.Scanf("%d", &a)
	// fmt.Scanf("%d", &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
	//
	// WRITE YOUR CODE HERE
	var c int
	c = a
	a = b
	b = c

	//
	fmt.Printf("a = %d, b = %d\n", a, b)
}
