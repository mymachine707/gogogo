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
	var answer string // javobni answerga biriktirish
	var s []string
	var y int
	var e int // a katta b dan
	var c int // a kichik b dan
	// qaysi son kattaligini tekshirish kerak
	if len(a) > len(b) {
		// a katta b dan
		e += 1
	} else if len(a) < len(b) {
		// a kichik b dan
		c += 1
	} else if len(a) == len(b) {
		// a teng b ga
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				// a katta b dan
				e += 1
				break
			} else if a[i] < b[i] {
				// a kichik b dan
				c += 1
				break
			}

		}
		answer = "0"

	}

	// slicelarni bir biriga tenglashtirish jarayoni bu kod o'zgartirilmasin

	if e == 1 { // a katta b dan
		y = len(a) - len(b)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, b...)
		b = s
		s = nil

	} else if c == 1 { // a kichik b dan
		y = len(b) - len(a)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, a...)
		a = s
		s = nil
	}

	var k []int
	if e == 1 { // a>b
		answer = ""
		for i := len(a); i > 0; i-- {
			if strconvInt(a[i-1]) < strconvInt(b[i-1]) {
				//

				if strconvInt(a[i-1]) == 0 {
					a[i-1] = "9"
					if i-2 >= 0 { // 0 ga teng bo'vb qosa
						if a[i-2] == "0" {
							a[i-2]=
						} else {
							a[i-2] = intconvstr(strconvInt(a[i-2]) - 1)
						}
					}

				} else {
					a[i-1] = intconvstr(strconvInt(a[i-1]) - 1)
					if i-2 >= 0 {
						a[i-2] = intconvstr(strconvInt(a[i-2]) - 1)
					}
				}
			}

			h := strconvInt(a[i]) - strconvInt(b[i])
			k = append(k, h)
		}
		fmt.Println("\tnumber1 > number2")
		fmt.Printf("\tAnswer number: %d\n", k)
	} else if c == 1 { // a<b
		answer = ""
		for i := 0; i < len(a); i++ {
			h := strconvInt(b[i]) - strconvInt(a[i])
			k = append(k, h)
		}
		fmt.Println("\tnumber1 < number2")
		fmt.Printf("\tAnswer number: %d\n", k)
	} else {
		fmt.Println("\tnumber1 = number2")
		fmt.Printf("\t   Answer number: %s\n", answer)
	}

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
