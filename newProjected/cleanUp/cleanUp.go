package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	// string...
	var num string

	fmt.Printf("%s", "\tEntered number 1: ")
	fmt.Scanf("%s", &num)
	//
	a := strings.Split(num, "")
	//
	var number bool
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

	for i := 1; i < len(a); i++ {
		if isInt(a[i]) {
			number = true
		} else {
			panic("This string is not a number!")
		}
	}

	var b []string
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == "0" {
			count++
		}
	}
	b = a[count:]

	var answer string

	for i := 0; i < len(b); i++ {
		answer += b[i]
	}

	fmt.Println("This number is negative = ", negative)
	fmt.Println("The first index of string is number = ", answer)

}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func str_Int(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		return 0 // sho'tda tanish bilishchili bo'b ketti err bersa 0 qayatrmasligi kere boshqa javob so'ra
	}
	return intVar
}
