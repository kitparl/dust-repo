package main

import "fmt"

func main(){
  numbers := [6]int{1,2,3,4,5,6}

// Access Elements of an Array
fmt.Println(numbers[1])

// Change elements of an array
numbers[1] = 10
fmt.Println(numbers[1])

// Array Initialization
numbers2 := [5]int{} // Not initialized
numbers3 := [5]int{1,2} // Partially initialized
numbers4 := [5]int{1,2,3,4,5} // fully initialized
fmt.Println(numbers2)
fmt.Println(numbers3)
fmt.Println(numbers4)

// Initialize only specific elements

numbers5 := [5]int{1:2, 3:4}
fmt.Println(numbers5)

// Length of an array

fmt.Println(len(numbers))

}
