package main

import "fmt"

func main() {
  numbers := []int{1,2,3,4,5,6,7,8,9,10}

// Access Elements of a Slice
fmt.Println(numbers[1])

// Change Elements of a Slice
numbers[2] = 50

fmt.Println(numbers[2])

// Append Elements To a Slice
// slice_name = append(slice_name, element1, element2, ...)

numbers = append(numbers, 100, 200, 300)
fmt.Println(numbers)

// Append One Slice To Another Slice
// slice3 = append(slice1, slice2...)

  myslice1 := []int{1,2,3}
  myslice2 := []int{4,5,6}
  myslice3 := append(myslice1, myslice2...)

  fmt.Println(myslice3)

// change the length of the slice

  arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
  my_slice1 := arr1[1:5] // Slice array
  fmt.Println(my_slice1)
  my_slice1 = arr1[1:3]
  fmt.Println(my_slice1)

// NOTE: If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
// copy(dest, src)

  examples := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  fmt.Println(examples)
  neededNumbers := examples[:len(examples)-10]
  fmt.Println(neededNumbers)
  numbersCopy := make([]int, len(neededNumbers))
  copy(numbersCopy, neededNumbers)
  fmt.Println(numbersCopy)

}
