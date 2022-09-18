package main

import (
	"strconv"
	"strings"
)

func main() {
	var a string = "1234"
	var b string = "1234"
	var k []string


	var e int // a katta b dan
	var c int // a kichik b dan
	// qaysi son kattaligini tekshirish kerak
	if len(a) > len(b) {
		// a katta b dan
		e += 1
	} else if len(a) < len(b) {
		// a kichik b dan
		c += 1

	for i := 0; i < count; i++ {

	}
	//fmt.Println(mult(num, b))
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

func mult(w, b string) []string {

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
	return l
}

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
