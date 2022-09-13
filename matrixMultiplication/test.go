package main

import "fmt"

func main() {
	arr1 := [2][2]float64{{1, 2}, {3, 4}}
	arr2 := [2][2]float64{{5, 6}, {7, 8}}

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
}
