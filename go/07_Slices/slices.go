package main

import "fmt"

// Slices are similar to arrays, but are more powerful and flexible.
// Like arrays, slices are also used to store multiple values of the same type in a single variable.
// However, unlike arrays, the length of a slice can grow and shrink as see fit.

func main(){

// In Go, there are several ways to create slices:



// 1. Using the []datatype{values} format

// slice_name := []datatype{values}

myslice := []int{1,2,3,5}

fmt.Println(myslice)
fmt.Println(len(myslice)) // len() is used to find the length of a slice
fmt.Println(cap(myslice)) // cap() it returns the capacity of a slice

// 2. Create a Slice With The make() Function

// slice_name := make([]type, length, capacity)
// Note: If the capacity parameter is not defined, it will be equal to length.

 myslice1 := make([]int, 5, 10)
 fmt.Println(myslice1)
 fmt.Println(len(myslice1))
 fmt.Println(cap(myslice1))


}
