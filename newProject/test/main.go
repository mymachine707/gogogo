package main

import "strconv"

func main() {

}

func fact(slice []string, index int) string {
	if index-3==0 || index-3 < 0{
		slice[0]=intconvstr(strconvInt(slice[0]-1))
		return "9"
	}
	slice[index-2] = "9"
		if index-3 > 0 { // 0 ga teng bo'vb qosa
			slice[index-3] = intconvstr(strconvInt(slice[index-3]) - 1)
	
	return fact(index-1)
}

/*
a-2 olishim kere
uni 9 ga tengliman
index-2>=0



*/
