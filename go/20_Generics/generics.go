package main

import "fmt"

// Generics allows us to write flexible and reusable code that can work with any data type, without sacrificing type safety.

func Print[T any](value T) { // Define the generic function fmt.Println(value)
}
// T is a type parameter
// The any cconstraint means that T can be any type


func Sum[T int | float64](a,b T) T {
  return a + b
}

// Defining Generic Types

type Pair[T any] struct {
  First, Second T
}

func main() {

  Print(42)
  Print("Hello")


  fmt.Println(Sum(1,3))
  fmt.Println(Sum(1.5,3.3))
  
  intPair := Pair[int]{First: 1, Second: 2}
  floatPair := Pair[float64]{First: 1.5, Second: 2.5}

  fmt.Println(intPair)
  fmt.Println(floatPair)

  

}


// Key Benefits of Generics in GO
// 1. Code Reusability
// 2. Type Safety
// 3. Improved Readability
