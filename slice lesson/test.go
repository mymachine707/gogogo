package main

import "fmt"

func main() {
	var star string
	var count int

	fmt.Printf("%s", "Enter count stars: ")
	fmt.Scanf("%d", &count)

	// romb tepasi

	for i := count; i > 0; i-- {
		var c int = count - i + 1

		for o := 0; o < i; o++ {
			star += " "
		}

		for l := 0; l < c; l++ {
			star += "*"
		}

		// yoni

		var s int = count - i

		for p := 0; p < s; p++ {
			star += "*"
		}
		fmt.Printf(star)
		fmt.Println()
		star = ""
	}

	// romb pasi
	for i := 0; i < count; i++ {

		var h int = count - i

		for p := 0; p < i; p++ {
			star += " "
		}

		for m := 0; m < h; m++ {
			star += "*"
		}

		var f int = count - i - 1

		for i := f; i > 0; i-- {
			star += "*"
		}
		fmt.Printf(star)
		fmt.Println()
		star = ""
	}
}
