package main

import "fmt"

func main() {
	
  var w, x int = 1, 2 // can declare multiple variables
	fmt.Println(w, x)
	x, w = w, x
	fmt.Println(w, x)


	// Using third variable

	a := 1
	b := 2
	temp := 0

	temp = a
	a = b
	b = temp

	fmt.Println(a, b)

	// using without third variable
	usingXor(a , b)

	// using function
	a, b = usingFunc(a , b)
	fmt.Println(a, b)
	
	// using arithmatic
	usingArthmatic(a, b)
}

func usingXor(a, b int) int {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a , b)
	return 0
}

func usingFunc(a, b int)(int, int){
	return b, a
}

func usingArthmatic(a, b int){
	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a , b)
}
