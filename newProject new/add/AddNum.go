package add

import (
	"strconv"
	"strings"
)

/*
func main() {

	fmt.Println("Sonlarni string holatida qo'shish")
	var num1 string = "546000000000000000000000000000000005"  // o'n ming
	var num2 string = "1200000000000000000000000000000000031" // o'n beshming

	fmt.Printf("%s", "\tEntered number 1: ")
	fmt.Scanf("%s", &num1)

	fmt.Printf("%s", "\tEntered number 2: ")
	fmt.Scanf("%s", &num2)

	fmt.Printf("\t     answer=%s\t\n", AddNum(num1, num2))

}*/

func add(num1, num2 string) string {
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
		z := str_Int(a[i]) + str_Int(b[i])
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
		answer += int_str(k[i])
	}

	return answer
} //------------------------------------------------main yopildi

// stringdan intga o'tqizadi
func str_Int(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		return 0 // sho'tda tanish bilishchili bo'b ketti err bersa 0 qayatrmasligi kere boshqa javob so'ra
	}
	return intVar
}

// intdan stringga o'tkizadi

func int_str(num int) string {
	str := strconv.Itoa(num)
	return str
}
