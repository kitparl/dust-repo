package main

import "fmt"

//  the defer keyword is used to delay the execution of a function until the surrounding function returns.
//  Deferred function calls are executed in LIFO (Last In, First Out) order, meaning the last deferred call will be executed first when the surrounding function ends.
func main() {
  fmt.Println("Start")
  // Defer this function call
  defer fmt.Println("End")

  fmt.Println("Middle")




  defer fmt.Println("First")
  defer fmt.Println("Second")
  defer fmt.Println("Third")

  fmt.Println("Main function executed")
}
