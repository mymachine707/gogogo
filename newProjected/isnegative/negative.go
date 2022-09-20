package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	// string...
	var num string

	fmt.Printf("%s", "\tEntered number 1: ")
	fmt.Scanf("%s", &num)

	fmt.Println(isnegatives(num))
}

func isnegatives(num string) bool {
	a := strings.Split(num, "")
	//
	var negative bool

	if a[0] == "+" {
		negative = false
	} else if a[0] == "-" {
		negative = true
	} else if isInt(a[0]) {
		negative = false
	} else {
		panic("This string is not a number!")
	}

	return negative

}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
