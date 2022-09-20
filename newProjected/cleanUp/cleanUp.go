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
	//
	cleanUp(num)
}

func cleanUp(num string) {
	a := strings.Split(num, "")
	//

	if a[0] == "+" {
		a = a[1:]
	} else if a[0] == "-" {
		a = a[1:]
	} else if !isInt(a[0]) {
		panic("This string is not a number!")
	}

	for i := 1; i < len(a); i++ {
		if !isInt(a[i]) {
			panic("This string is not a number!")
		}
	}

	var b []string
	count := 0

	for i := 0; i < len(a); i++ {

		if a[i] == "0" {
			count++
		} else {
			break
		}
	}
	b = a[count:]

	var answer string

	for i := 0; i < len(b); i++ {
		answer += b[i]
	}

	fmt.Println(answer)
} //-end function

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
