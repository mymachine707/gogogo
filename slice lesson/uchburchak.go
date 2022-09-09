package main

import "fmt"

func main() {
	var star string
	var count int
	fmt.Printf("%s", "Enter count stars: ")
	fmt.Scanf("%d", &count)

	// To'gri burchakli uchburchak burchak chap tomon pastda
	for i := 0; i < count; i++ {
		for j := 0; j < i; j++ {
			star += "*"
		}
		fmt.Printf("%s", star)
		fmt.Println()
		star = ""
	}

	// To'gri burchakli uchburchak burchak chap tomon tepada
	for i := count; i > 0; i-- {
		for j := 0; j < i; j++ {
			star += "*"
		}
		fmt.Printf("%s", star)
		fmt.Println()
		star = ""
	}

	//To'gri burchakli uchburchak burchak tepa tomon o'ngda

	for i := count; i > 0; i-- {
		var c int = count
		for k := 0; k < i; k++ {
			star += " "
			c--
		}
		for l := 0; l < c; l++ {
			star += "*"
		}
		fmt.Printf("%s", star)
		fmt.Println()
		star = ""

	}

	//To'gri burchakli uchburchak burchak tepa
	for i := 0; i < count; i++ {
		var c int = count
		for k := 0; k < i; k++ {
			star += " "
			c--
		}
		for l := 0; l < c; l++ {
			star += "*"
		}
		fmt.Printf("%s", star)
		fmt.Println()
		star = ""

	}

}
