package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Sonlarni string holatida qo'shish")
	var num1 string = "5465" // o'n ming
	var num2 string = "0001" // o'n beshming

	fmt.Printf("%s", "\tEntered number 1: ")
	fmt.Scanf("%s", &num1)

	fmt.Printf("%s", "\tEntered number 2: ")
	fmt.Scanf("%s", &num2)

	a := strings.Split(num1, "")
	b := strings.Split(num2, "")

	// kiruvchi stringlarni uzunligini tenglashtirish======================

	var s []string
	var y int

	if len(a) > len(b) {
		y = len(a) - len(b)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, b...)
		b = s
		s = nil
	} else {
		y = len(b) - len(a)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, a...)
		a = s
		s = nil
	}

	//stringlarni slice holatida har bir elementini qo'shib yangi slice hosil qilish

	var answer string // javobni ansverga biriktirish

	for i := 0; i < len(a); i++ {
		answer += intconvstr(strconvInt(a[i]) + strconvInt(b[i]))
	}

	fmt.Printf("\tAnswer number: %s\n", answer)
} //------------------------------------------------main yopildi

// stringdan intga o'tqizadi
func strconvInt(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		return 0 // sho'tda tanish bilishchili bo'b ketti err bersa 0 qayatrmasligi kere boshqa javob so'ra
	}
	return intVar
}

// intdan stringga o'tkizadi

func intconvstr(num int) string {
	str := strconv.Itoa(num)
	return str
}
