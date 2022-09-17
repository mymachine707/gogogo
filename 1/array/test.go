package main

import "fmt"

func main() {
	A := [2][2]int{{1, 2}, {3, 4}}
	B := [2][2]int{{5, 6}, {7, 8}}
	for i := 0; i < len(A); i++ {
		for s := 0; s < len(B); s++ {
			fmt.Printf("A[%d][%d]=%d\n", i, s, A[i][s])

		}
	}

	for i := 0; i < len(A); i++ {
		for s := 0; s < len(B); s++ {

			fmt.Printf("B[%d][%d]=%d\n", i, s, B[i][s])
		}
	}

	var C [][]int

	C = make([][]int, 2)
	for i := 0; i < 2; i++ {
		C[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			for k := 0; k < len(A)+1; k++ {
				C[i][j] = A[i][k] * B[k][j]
			}
		}

		fmt.Println(C)
	}
}
