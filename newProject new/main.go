package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num1 string
	var num2 string

	fmt.Printf("%s", "\tEntered number 1: ")
	fmt.Scanf("%s", &num1)

	fmt.Printf("%s", "\tEntered number 2: ")
	fmt.Scanf("%s", &num2)
	p := Add(num1, num2)
	q := minus(num1, num2)
	fmt.Printf("answer to nomber add=%s\n", p)
	fmt.Printf("answer to nomber minus=%s\n", q)
}

//====================================================

//======================================================================== musbat sonlarni qo'shish

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

//==================================================== musbat sonlarni ayirish

func minus(num1, num2 string) string {
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
	var z []string

	for i := 0; i < len(a); i++ {
		z = append(z, a[i])
	}

	//a bilan b berilgan slice holatida
	if e == 1 {
		fmt.Println("\nnumber1 > number 2")
		// a > b bo'lsa:
		for i := len(a); i > 0; i-- {
			if str_Int(a[i-1])-str_Int(b[i-1]) > 0 { // Agar joriy indekslar ayirmasi musbat bo'lsa
				z[i-1] = int_str(str_Int(a[i-1]) - str_Int(b[i-1])) // z joriy indeksiga ayirma qiymati beriladi
			} else if str_Int(a[i-1])-str_Int(b[i-1]) == 0 { // Agar joriy indekslar ayirmasi 0 ga teng bo'lsa
				z[i-1] = int_str(0) // z joriy indeksiga 0 qiymat beriladi
			} else if str_Int(a[i-1])-str_Int(b[i-1]) < 0 { // agar joriy indekslar ayirmasi 0 dan kichik bo'lsa
				a[i-1] = int_str(str_Int(a[i-1]) + 10)              // katt sonning joriy indeksiga 10 soni qo'shiladi va qaytadan yenglashtiriladi
				z[i-1] = int_str(str_Int(a[i-1]) - str_Int(b[i-1])) // z ning joriy qiymatini sonlarning joriy qiymatlari ayirmasiga tenglashtiriladi va...
				if str_Int(a[i-2]) > 0 {                            // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 dan katta bo'lsa
					a[i-2] = int_str(str_Int(a[i-2]) - 1) // katta sonning joriy qiymatidan bitta oldingi sonnini bittaga kamaytiriladi
				} else if str_Int(a[i-2]) == 0 { // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 ga teng bo'lsa...
					a[i-2] = "9"           // katta sonning joriy qiymatidan bitta oldingi sonnini 9 soniga tenglashtiriladi
					var boolen bool = true // for loop uchun boolen qiyamtini true holatida ochildi
					var m int = 3          // indekslarni tekshirish uchun m int shaklida ochildi va katta sonning joriy qiymatidan ikkita oldingi qiymatini tekshirish uchun m sonini 3 ga tenglandi
					for boolen {           // for loop ochildi
						if str_Int(a[i-m]) > 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 dan katta bo'lsa
							a[i-m] = int_str(str_Int(a[i-m]) - 1) // katta sonning joriy indeksidan ikkita oldingi indeksidan 1 soni ayiriladi
							boolen = false                        // boolen false qiymatini oladi
							break                                 // for loopdan chiqiladi
						} else if str_Int(a[i-m]) == 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 ga teng bo'lsa
							a[i-m] = "9" // Katta sonning joriy indeksidan ikkita oldingi indeksi "9" ga tenglanadi
							m++          // m soni 1 ga oshiriladi va for loop qaytadan tekshiruvni boshlaydi
						}

					}
				}
			}

		}

	} else if c == 1 {
		fmt.Println("\nnumber1 < number 2")
		// a < b bo'lsa:
		for i := len(b); i > 0; i-- {
			if str_Int(b[i-1])-str_Int(a[i-1]) > 0 { // Agar joriy indekslar ayirmasi musbat bo'lsa
				z[i-1] = int_str(str_Int(b[i-1]) - str_Int(a[i-1])) // z joriy indeksiga ayirma qiymati beriladi
			} else if str_Int(b[i-1])-str_Int(a[i-1]) == 0 { // Agar joriy indekslar ayirmasi 0 ga teng bo'lsa
				z[i-1] = int_str(0) // z joriy indeksiga 0 qiymat beriladi
			} else if str_Int(b[i-1])-str_Int(a[i-1]) < 0 { // agar joriy indekslar ayirmasi 0 dan kichik bo'lsa
				b[i-1] = int_str(str_Int(b[i-1]) + 10)              // katt sonning joriy indeksiga 10 soni qo'shiladi va qaytadan yenglashtiriladi
				z[i-1] = int_str(str_Int(b[i-1]) - str_Int(a[i-1])) // z ning joriy qiymatini sonlarning joriy qiymatlari ayirmasiga tenglashtiriladi va...
				if str_Int(b[i-2]) > 0 {                            // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 dan katta bo'lsa
					b[i-2] = int_str(str_Int(b[i-2]) - 1) // katta sonning joriy qiymatidan bitta oldingi sonnini bittaga kamaytiriladi
				} else if str_Int(b[i-2]) == 0 { // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 ga teng bo'lsa...
					b[i-2] = "9"           // katta sonning joriy qiymatidan bitta oldingi sonnini 9 soniga tenglashtiriladi
					var boolen bool = true // for loop uchun boolen qiyamtini true holatida ochildi
					var m int = 3          // indekslarni tekshirish uchun m int shaklida ochildi va katta sonning joriy qiymatidan ikkita oldingi qiymatini tekshirish uchun m sonini 3 ga tenglandi
					for boolen {           // for loop ochildi
						if str_Int(b[i-m]) > 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 dan katta bo'lsa
							b[i-m] = int_str(str_Int(b[i-m]) - 1) // katta sonning joriy indeksidan ikkita oldingi indeksidan 1 soni ayiriladi
							boolen = false                        // boolen false qiymatini oladi
							break                                 // for loopdan chiqiladi
						} else if str_Int(b[i-m]) == 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 ga teng bo'lsa
							b[i-m] = "9" // Katta sonning joriy indeksidan ikkita oldingi indeksi "9" ga tenglanadi
							m++          // m soni 1 ga oshiriladi va for loop qaytadan tekshiruvni boshlaydi
						}

					}
				}
			}

		}
	}
	//natijani birlashtirish
	var answers string
	for i := 0; i < len(z); i++ {
		answers += z[i]
	}

	//==========================
	return answers
} //------------------------- manfiy sonlarni ayirish


func minus(num1, num2 string) string {
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
	var z []string

	for i := 0; i < len(a); i++ {
		z = append(z, a[i])
	}

	//a bilan b berilgan slice holatida
	if e == 1 {
		fmt.Println("\nnumber1 > number 2")
		// a > b bo'lsa:
		for i := len(a); i > 0; i-- {
			if str_Int(a[i-1])-str_Int(b[i-1]) > 0 { // Agar joriy indekslar ayirmasi musbat bo'lsa
				z[i-1] = int_str(str_Int(a[i-1]) - str_Int(b[i-1])) // z joriy indeksiga ayirma qiymati beriladi
			} else if str_Int(a[i-1])-str_Int(b[i-1]) == 0 { // Agar joriy indekslar ayirmasi 0 ga teng bo'lsa
				z[i-1] = int_str(0) // z joriy indeksiga 0 qiymat beriladi
			} else if str_Int(a[i-1])-str_Int(b[i-1]) < 0 { // agar joriy indekslar ayirmasi 0 dan kichik bo'lsa
				a[i-1] = int_str(str_Int(a[i-1]) + 10)              // katt sonning joriy indeksiga 10 soni qo'shiladi va qaytadan yenglashtiriladi
				z[i-1] = int_str(str_Int(a[i-1]) - str_Int(b[i-1])) // z ning joriy qiymatini sonlarning joriy qiymatlari ayirmasiga tenglashtiriladi va...
				if str_Int(a[i-2]) > 0 {                            // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 dan katta bo'lsa
					a[i-2] = int_str(str_Int(a[i-2]) - 1) // katta sonning joriy qiymatidan bitta oldingi sonnini bittaga kamaytiriladi
				} else if str_Int(a[i-2]) == 0 { // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 ga teng bo'lsa...
					a[i-2] = "9"           // katta sonning joriy qiymatidan bitta oldingi sonnini 9 soniga tenglashtiriladi
					var boolen bool = true // for loop uchun boolen qiyamtini true holatida ochildi
					var m int = 3          // indekslarni tekshirish uchun m int shaklida ochildi va katta sonning joriy qiymatidan ikkita oldingi qiymatini tekshirish uchun m sonini 3 ga tenglandi
					for boolen {           // for loop ochildi
						if str_Int(a[i-m]) > 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 dan katta bo'lsa
							a[i-m] = int_str(str_Int(a[i-m]) - 1) // katta sonning joriy indeksidan ikkita oldingi indeksidan 1 soni ayiriladi
							boolen = false                        // boolen false qiymatini oladi
							break                                 // for loopdan chiqiladi
						} else if str_Int(a[i-m]) == 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 ga teng bo'lsa
							a[i-m] = "9" // Katta sonning joriy indeksidan ikkita oldingi indeksi "9" ga tenglanadi
							m++          // m soni 1 ga oshiriladi va for loop qaytadan tekshiruvni boshlaydi
						}

					}
				}
			}

		}

	} else if c == 1 {
		fmt.Println("\nnumber1 < number 2")
		// a < b bo'lsa:
		for i := len(b); i > 0; i-- {
			if str_Int(b[i-1])-str_Int(a[i-1]) > 0 { // Agar joriy indekslar ayirmasi musbat bo'lsa
				z[i-1] = int_str(str_Int(b[i-1]) - str_Int(a[i-1])) // z joriy indeksiga ayirma qiymati beriladi
			} else if str_Int(b[i-1])-str_Int(a[i-1]) == 0 { // Agar joriy indekslar ayirmasi 0 ga teng bo'lsa
				z[i-1] = int_str(0) // z joriy indeksiga 0 qiymat beriladi
			} else if str_Int(b[i-1])-str_Int(a[i-1]) < 0 { // agar joriy indekslar ayirmasi 0 dan kichik bo'lsa
				b[i-1] = int_str(str_Int(b[i-1]) + 10)              // katt sonning joriy indeksiga 10 soni qo'shiladi va qaytadan yenglashtiriladi
				z[i-1] = int_str(str_Int(b[i-1]) - str_Int(a[i-1])) // z ning joriy qiymatini sonlarning joriy qiymatlari ayirmasiga tenglashtiriladi va...
				if str_Int(b[i-2]) > 0 {                            // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 dan katta bo'lsa
					b[i-2] = int_str(str_Int(b[i-2]) - 1) // katta sonning joriy qiymatidan bitta oldingi sonnini bittaga kamaytiriladi
				} else if str_Int(b[i-2]) == 0 { // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 ga teng bo'lsa...
					b[i-2] = "9"           // katta sonning joriy qiymatidan bitta oldingi sonnini 9 soniga tenglashtiriladi
					var boolen bool = true // for loop uchun boolen qiyamtini true holatida ochildi
					var m int = 3          // indekslarni tekshirish uchun m int shaklida ochildi va katta sonning joriy qiymatidan ikkita oldingi qiymatini tekshirish uchun m sonini 3 ga tenglandi
					for boolen {           // for loop ochildi
						if str_Int(b[i-m]) > 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 dan katta bo'lsa
							b[i-m] = int_str(str_Int(b[i-m]) - 1) // katta sonning joriy indeksidan ikkita oldingi indeksidan 1 soni ayiriladi
							boolen = false                        // boolen false qiymatini oladi
							break                                 // for loopdan chiqiladi
						} else if str_Int(b[i-m]) == 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 ga teng bo'lsa
							b[i-m] = "9" // Katta sonning joriy indeksidan ikkita oldingi indeksi "9" ga tenglanadi
							m++          // m soni 1 ga oshiriladi va for loop qaytadan tekshiruvni boshlaydi
						}

					}
				}
			}

		}
	}
	//natijani birlashtirish
	var answers string
	for i := 0; i < len(z); i++ {
		answers += z[i]
	}

	//==========================
	return answers
} //-------------------------
