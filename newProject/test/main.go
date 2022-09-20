fmt.Println("number1 < number 2")
// a > a bo'lsa:
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