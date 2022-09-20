// Golang program
// using unicode.IsNumber() Function

package main

import (
	"fmt"
	"unicode"
)

func main() {
	// declaring a variable mystr
	myStr := "89797987"
	// a for loop to iterate the variable myStr
	for _, char := range myStr {
		// the if statement checks if each character passed is a number
		// The displayed result depends on the return value
		if unicode.IsNumber(char) {
			fmt.Println(string(char), char, "is number rune")
		} else {
			fmt.Println(string(char), char, "is not a number rune")
		}
	}
}
