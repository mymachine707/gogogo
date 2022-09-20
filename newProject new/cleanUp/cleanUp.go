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
	var negative bool
	var number bool
	//

	var first string = a[0]
	//
	//
	if first == "+" {
		negative = false
	} else if first == "-" {
		negative = true
	} else if unicode.IsNumber(first) {
		number = true
		//
		fmt.Println("This string is number = ", number)
		//
		negative = false
	} else {
		panic("This string is not a number")
	}

	fmt.Println("Number is negative = ", negative)

	//
	//
	for i := 1; i < len(a); i++ {

	}

}

func str_Int(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		return 0 // sho'tda tanish bilishchili bo'b ketti err bersa 0 qayatrmasligi kere boshqa javob so'ra
	}
	return intVar
}
