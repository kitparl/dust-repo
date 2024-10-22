package main

import "fmt"


// Go has thee functions to output
// 1. Print()
// 2. Println()
// 3. Printf()
func main(){


  var i,j string = "Hello","Golang"
  fmt.Print(i,j)
  // Print() inserts a space between the arguments if neither are strings:


  fmt.Println(i,j)
  // The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:


  // %v is used to print the value of the arguments
  // %T is used to print the type of the arguments

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)

  /*
  General Formatting Verbs

  Verb	Description
  %v	Prints the value in the default format
  %#v	Prints the value in Go-syntax format
  %T	Prints the type of the value
  %%	Prints the % sign
  */


}
