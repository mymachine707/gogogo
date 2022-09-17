package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var a int
	var b int
	var c int

	return func() int {

		if c < 2 {
			if c == 1 {
				c++
				return 1
			}
			c++
			return 0
		}

		a, b = b, a
		b = a + b // 1, 2
		c++

		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
