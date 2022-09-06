package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(num int) {
	//
	// WRITE YOUR CODE HERE
	if num%3 == 0 && num%5 == 0 {
		fmt.Printf("FizzBuzz\n")
	} else if num%3 == 0 {
		fmt.Printf("Fizz\n")
	} else if num%5 == 0 {
		fmt.Printf("Buzz\n")
	} else {
		fmt.Println(num)
	}

	//
}
