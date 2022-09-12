package main

import "fmt"

func main() {
	package main

	import "fmt"
	
	func main() {
	  arr1 := [2][2]float64{{1, 2}, {3, 4}}
	  arr2 := [2][2]float64{{5, 6}, {7, 8}}
	
	  var ans [2][2]float64
	
	  for i := 0; i < len(arr1); i++ {
		for s := 0; s < len(arr1); s++ {
		  for p := 0; p < len(arr1); p++ {
			if ans[i][s] == 0 {
			  ans[i][s] = arr1[i][p] * arr2[p][s]
			  fmt.Println("*")
			}
			ans[i][s] += arr1[i][p] * arr2[p][s]
			fmt.Println("+")
	
		  }
		}
	  }
	
	  fmt.Println(ans)
	}
	
}
