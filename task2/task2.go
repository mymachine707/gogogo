package main

import "fmt"

func main() {
	var a, b int = 6, 4

	// fmt.Scanf("%d", &a)
	// fmt.Scanf("%d", &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Println(a, "is odd (toq son): ", isOdd(a))
	fmt.Println(b, "is even (juft son): ", isEven(b))
}

func isEven(num int) bool {
	//
	// WRITE YOUR CODE HERE
	if num%2 == 0 {
		return true
	}
	return false
	//
}

func isOdd(num int) bool {
	//
	// WRITE YOUR CODE HERE
	if num%2 != 0 {
		return true
	}
	return false
	//
}
