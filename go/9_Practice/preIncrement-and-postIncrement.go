package main

import "fmt"

func main() {
	// Go does not support pre-increment (++x) at all,
	// Only post-increment and post-decrement are supported
	
	// y = ++x; // C-style

	var x int = 0
	x += 1      // Increment first
	y := x      // Then use the value

	fmt.Println(y)


	// y := x++ // ❌ compile-time error in Go


	a := 0
	a++
  fmt.Println(a)


}
