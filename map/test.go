package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["new"] = 1
	m["new1"] = 11
	m["new"]++
	fmt.Println(m)
}
