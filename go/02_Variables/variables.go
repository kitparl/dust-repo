package main

import "fmt"

func main() {


// 1. With the var keyword:
  var v = "i am a" // var declares 1 or more variables
  fmt.Println(v)


  // Go multiple variable declaration
  var w, x int = 1, 2 // can declare multiple variables
  fmt.Println(w,x)


  var d = true
  fmt.Println(d)


  var e int     // defalut value is 0
  fmt.Println(e)

// 2. With the := sign:
  f := "anything" // The := syntax is shorthand for declaring and assigning a variable. This syntax only avaiable inside functions
  fmt.Println(f)

  num1, num2 := 1, 2 // Multple declaration
  fmt.Println(num1, num2)


  /*

   Rules for Using :=

   1. Inside Functions : The := operator can only be used inside functions (or method bodies). You cannot use it at the package level (outside functions).
   2. Redeclaration : If a variable has already been declared in the same scope, you cannot use := to assign a new value to it unless at least one of the variables in the declaration is new.


  y := 30 // Valid if y was previously declared using var in the same scope
  y = 20
  fmt.Println(y)


- Use := only for declaring and initializing new variables.

- For reinitializing or updating existing variables, use =.

- You can use := again only if at least one of the variables in the statement is new.

  */

  var a string
  var b int
  var c bool

  fmt.Println(a) // default value is ""
  fmt.Println(b) // default value is 0 
  fmt.Println(c) // default value is false


  /*
  * var
  * 1. Can used inside and outside functions
  * 2. Variables declaration and value assignment can be done separately
  *
  *
  * :=
  * 1. Can only used inside functions
  * 2. Variables declaration and value assignment can be done together
  */


  // ### Go variable declaration in a Block

  var (
    rollNo int
    age int = 1
    name string = "pranshu"
  )

  fmt.Println(rollNo, age, name)

}
