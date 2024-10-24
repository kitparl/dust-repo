package main

import "fmt"

// Pointers are variables that hold the memory address of another variable
// They are used to pass references to values rather than the values themselves
func main() {
  // Pointers declaration
  var pointer *int
  fmt.Println(pointer)
  
  // getting the address of a variable
  var num = 10
  pointer = &num  // The &i syntax gives the memory address of num
  fmt.Println(pointer)


  // Dereferencing a pointer (*) --->>> current value holded at that address
  fmt.Println(*pointer)


  // Modifing varaibles using pointers
  *pointer = 403
  fmt.Println(num)
}
