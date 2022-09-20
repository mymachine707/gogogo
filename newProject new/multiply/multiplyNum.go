package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Sonlarni string holatida ko'paytirish")
	var a string // o'n ming
	var b string // o'n beshming

	fmt.Printf("%s", "\tEntered number 1: ")
	fmt.Scanf("%s", &a)

	fmt.Printf("%s", "\tEntered number 2: ")
	fmt.Scanf("%s", &b)

	var k []string

	var e int // a katta b dan
	var c int // a kichik b dan
	// qaysi son kattaligini tekshirish kerak
	if len(a) >= len(b) {
		// a katta b dan
		e += 1

	} else if len(a) < len(b) {
		// a kichik b dan
		c += 1

	}
	var ans string
	if e == 1 {
		f := strings.Split(b, "")
		for i := len(f); i > 0; i-- {
			ans = mult(a, f[i-1])
			for l := 0; l < len(f)-i; l++ {
				ans += "0"
			}

			k = append(k, ans)
		}

	} else if c == 1 {
		f := strings.Split(a, "")
		for i := len(f); i > 0; i-- {
			ans = mult(b, f[i-1])
			for l := 0; l < len(f)-i; l++ {
				ans += "0"
			}
			k = append(k, ans)

		}

	}

	var answers string

	for i := 0; i < len(k); i++ {
		answers = adder(answers, k[i])
	}
	fmt.Println("\t   Answer number:", answers)
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// katta xonali sonni bir xonali songa ko'paytirish

func mult(w, b string) string {

	a := strings.Split(w, "")

	var l []string = a

	var p int
	var m int = 0
	for i := len(a); i > 0; i-- {
		p = strconvInt(a[i-1]) * strconvInt(b)
		if i-1 == 0 {
			p = strconvInt(a[i-1]) * strconvInt(b)
			p += m
			m = 0
			l[i-1] = intconvstr(p)
		} else if p < 10 {
			if p+m >= 10 {
				p += m
				m = 0
				q := strings.Split(intconvstr(p+m), "")
				l[i-1] = q[1]
				m = strconvInt(q[0])
			} else {
				l[i-1] = intconvstr(p)
			}
		} else if p >= 10 {
			p += m
			m = 0
			q := strings.Split(intconvstr(p), "")
			l[i-1] = q[1]
			m = strconvInt(q[0])
		}
	}
	answer := ""

	for i := 0; i < len(l); i++ {
		answer += l[i]
	}
	return answer
}

// katta xonali sonlarni bir biriga qo'shish
//============

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
