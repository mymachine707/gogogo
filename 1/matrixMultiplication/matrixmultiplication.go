package main

import "fmt"

func main() {

	var n int
	var m int

	fmt.Printf("%s", "n=")
	fmt.Scanf("%d", &n)

	fmt.Printf("%s", "m=")
	fmt.Scanf("%d", &m)

	fmt.Printf("n=%d\tm=%d\n", n, m)

	/*
		var A [][]float64
		var B [][]float64

		A = make([][]float64, n)
		B = make([][]float64, m)

		ans := make([][]float64, 2)
		for i := 0; i < len(arr1); i++ {
			ans[i] = make([]float64, 2)
			for s := 0; s < len(arr1); s++ {
				ans[i][s] = 0

				for p := 0; p < len(arr1); p++ {
					ans[i][s] += arr1[i][p] * arr2[p][s]
					fmt.Printf("ans[%d][%d]+=%0.0f*%0.0f\n", i, s, arr1[i][p], arr2[p][s])

				}
			}
		}

		fmt.Println(ans)

	*/
}
