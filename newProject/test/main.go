package main

import "fmt"

func main() {
	var m int
	var x bool = true
	for x {
		if m == 100 {
			x = false
		}
		fmt.Println(m)
		m++
	}
	fmt.Println("The end")
}
