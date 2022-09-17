package main

func main() {
	DisplayNumberInReverseOrderWithDefer()
}

func DisplayNumberInReverseOrderWithDefer() {
	for i := 0; i < 100000000; i++ {
		//
		// WRITE YOUR CODE HERE
		defer A()
		//
	}
}

func A() {
	a := "12"
	a = a + "2"
}
