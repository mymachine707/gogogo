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
	var first string = a[0]
	//
	for i := 1; i < len(a); i++ {
		if isInt(a[i]) {
			number = true
		} else {
			number = false
			return number
		}
	}
	/*
		fmt.Println("This number is negative = ", negative)
		fmt.Println("The first index of string is number = ", number)
	*/
	//
	//

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
