package main

import "fmt"

func main() {

	// Standard Loop
	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

	// While style loop
	start, end := 0, 5
	for start<end {
		fmt.Println(start, end)
		start++
	}

	// Infinite Loop
  for {
    fmt.Println("Hello")
  }
}
