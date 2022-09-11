package main

import "fmt"

func main() {

	var A [][]float64

	var B [][]float64

	var C [][]float64

	var n int
	var m int

	fmt.Printf("%s", "n=")
	fmt.Scanf("%d", &n)

	fmt.Printf("%s", "m=")
	fmt.Scanf("%d", &m)

	fmt.Printf("n=%d\tm=%d\n", n, m)

	A = make([][]float64, n)
	for i := 0; i < n; i++ {
		A[i] = make([]float64, m)
		for s := 0; s < m; s++ {
			var in float64
			fmt.Printf("A[%d][%d]=", i, s)
			fmt.Scanf("%f", &in)
			A[i][s] = in

		}
	}

	B = make([][]float64, n)
	for i := 0; i < n; i++ {
		B[i] = make([]float64, m)
		for s := 0; s < m; s++ {
			var in float64
			fmt.Printf("B[%d][%d]=", i, s)
			fmt.Scanf("%f", &in)
			B[i][s] = in

		}
	}

	fmt.Printf("%s", "Array A:\n")
	for i := 0; i < len(A); i++ {
		fmt.Println(A[i])
	}

	fmt.Printf("%s", "Array B:\n")

	for i := 0; i < len(B); i++ {
		fmt.Println(B[i])
	}

	fmt.Printf("%s", "Array C:\n")

	for i := 0; i < len(C); i++ {
		fmt.Println(C[i])
	}
}
