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
	isnegatives(num)
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

func str_Int(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		return 0 // sho'tda tanish bilishchili bo'b ketti err bersa 0 qayatrmasligi kere boshqa javob so'ra
	}
	return intVar
}
