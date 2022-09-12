package main

import (
	"fmt"
	"strings"
)

// Needed to use Split

func main() {
	str := "akjkshljkdbnsjjhhgfksdafhnlk"
	split := strings.Split(str, "")
	m := make(map[string]int)

	fmt.Println(split)
	fmt.Println("The length of the slice is:", len(split))
	for i := 0; i < len(split); i++ {
		if IsLetter(split[i]) {
			m[split[i]]++
			//fmt.Println(split[i])
		}
	}
	fmt.Println(m)
}

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
