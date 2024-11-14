package main

import "fmt"

// Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

// In Go, there are two ways to declare an array:

func main(){

// 1. With the var keyword

/*
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred
*/
var arr1 = [3]int{1,2,3}
var arr2 = [...]int{1,2,3,8}
fmt.Println(arr1)
fmt.Println(arr2)



// 2. With the := sign:
/*
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred
*/
var arr3 = [5]int{4,6,1,2,3}
arr4 := [...]int{4,5,6,7,8}
fmt.Println(arr3)
fmt.Println(arr4)
}
