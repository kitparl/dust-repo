package main

import "fmt"

// A variadic function is a function that can take a variable number of arguments (similar to varargs in java)
func main() {
total := sum(1, 2, 3, 4, 5)
fmt.Println(total)

newSum := sum(1, 2, 3)
fmt.Println(newSum)
}


func sum(a ...int) int {
  sum := 0
  for _, v := range a {
    sum += v
  }
  return sum
}
