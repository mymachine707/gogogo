package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Sonlarni string holatida ayirish")
	var num1 string = "27564" // o'n ming
	var num2 string = "10987" // o'n beshming
	/*
		fmt.Printf("%s", "\tEntered number 1: ")
		fmt.Scanf("%s", &num1)

		fmt.Printf("%s", "\tEntered number 2: ")
		fmt.Scanf("%s", &num2)
	*/
	a := strings.Split(num1, "")
	b := strings.Split(num2, "")

	// kiruvchi stringlarni uzunligini tenglashtirish======================

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
	//==========================

	// asosiy shartlar
	var z []string = a
	var x bool = false
	var u int

	//a bilan b berilgan slice holatida

	if e == 1 { // a kotta b dan bo'gan holatida
		for i := len(a); i > 0; i-- {
			if x == false && strconvInt(a[i-1])-strconvInt(b[i-1]) > 0 {
				z[i-1] = intconvstr(strconvInt(a[i-1]) - strconvInt(b[i-1]))                          //z
				fmt.Printf("--->>>1--->: %s=%d-%d\n", z[i-1], strconvInt(a[i-1]), strconvInt(b[i-1])) //--->>>1--->:
			} else if x == false && strconvInt(a[i-1])-strconvInt(b[i-1]) < 0 { //------------------------------------------2

				u = strconvInt(a[i-1]) + 10
				z[i-1] = intconvstr(u - strconvInt(b[i-1]))

				fmt.Printf("--->>>2--->: %s=%d-%d\n", z[i-1], e, strconvInt(b[i-1])) //--->>>2--->:
				x = true
				//
			} else if x == true && strconvInt(a[i-1])-strconvInt(b[i-1]) > 0 {
				//
				if strconvInt(a[i-1])-1-strconvInt(b[i-1]) > 0 {
					z[i-1] = intconvstr(strconvInt(a[i-1]) - 1 - strconvInt(b[i-1]))
					fmt.Printf("--->>>3--->: %s=%d-%d\n", z[i-1], strconvInt(a[i-1]), strconvInt(b[i-1])) //--->>>3--->:
					x = false
				} else if strconvInt(a[i-1])-1-strconvInt(b[i-1]) < 0 {
					a[i-1] = intconvstr(strconvInt(a[i-1]) + 9)
					z[i-1] = intconvstr(strconvInt(a[i-1]) - strconvInt(b[i-1]))
					fmt.Printf("--->>>4--->: %s=%d-%d\n", z[i-1], strconvInt(a[i-1]), strconvInt(b[i-1])) //--->>>4--->:
					x = true
				}
			}
			fmt.Println(i - 1)
		}
	}

	// natijani birlashtirish
	//var answers string
	/*
		for i := 0; i < len(h); i++ {
			answers = adder(answers, h[i])
		}*/

	fmt.Println("\t    number1:", num1)
	fmt.Println("\t    number2:", num2)
	fmt.Println("\t   Answer number:", z)

	//==========================

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

// qo'shish funksiyasi
func adder(num1, num2 string) string {
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
	var k []int

	var m int
	for i := 0; i < len(a); i++ {
		z := strconvInt(a[i]) + strconvInt(b[i])
		k = append(k, z)

		for i := len(k); i > 0; i-- {
			if k[i-1] > 10 {
				if i-1 > 0 {
					m = k[i-1] - 10
					k[i-1] = m
					if i-2 >= 0 {
						k[i-2]++
					}
				} else if i-1 == 0 {
					break
				}

			}
		}
	}

	for i := 0; i < len(k); i++ {
		answer += intconvstr(k[i])
	}

	return answer
}

//============
