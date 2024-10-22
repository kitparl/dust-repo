package main

import "fmt"

// When a constant is declared, it is not possible to change the value later

const Pi = 3.14
// The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
func main() {
  const World = "Pranshu"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  // Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
  // Constants can be declared both inside and outside of a function
  

  // ### Types of constants ####
  // 1. Typed Constants

  const AGE int = 12
  fmt.Println(AGE)

  // 2. Untyped Constants

  const GENDER = "male"
  fmt.Println(GENDER)
  // In this case, the type of the constant is inferred from the value (means the compiler decides the type of the constant, based on the value).


// ### Multiple constants can be grouped together into a block for readability:
  
const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
