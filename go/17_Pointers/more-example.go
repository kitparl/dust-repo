package main

import "fmt"

func main() {
  var x int = 10;
  var ptr *int = &x // address in hexadecimal

  fmt.Println("x = ", x)
  fmt.Println("ptr = ", ptr)
  fmt.Println("value of x = ", *ptr)


  // ####### Memory Management with pointers
  // in go memory management is automatic and handled by garbage collector
  // Go also provides the ability to allocate memory dynamically using pointers.
  // Dynamic memory allocation allows us to create data structures that can grow and shrink as needed


  fmt.Println()
  var ptr2 *int = new(int) //  use the new function to allocate memory for an integer value.
  fmt.Println("ptr2 = ", ptr2)
  *ptr2 = 10
  fmt.Println("ptr2 = ", *ptr2)

  ptr2 = nil
  fmt.Println("ptr2 = ", ptr2)


  // ######## Pass by Value vs. Pass by Reference
  // all functions are pass by value
  // When we pass a value to a function, the value is copied to the function

  // We can use pointers to pass by reference
  // When we pass a variable to a function, the variable is passed by reference

  fmt.Println()
  pass := 10
  fmt.Println(pass)
  addOne(&pass)
  fmt.Println("addOne func", pass) 


  // #########  Pointer Arithmetic:
  // It is the ability to perform arithmetic operations on pointers, such as additon, substraction,
  // In Gom Pointer is limited to a few use cases.
  // For example, We can use pointer aithmetic to access of an array, but we cannot use it to perform arbitrary pointer arithmetic like in C/C++.

  fmt.Println()
  arr := [3]int{99, 20, 30}
  pointr := &arr[0]
  fmt.Println(*pointr)
  *pointr++
  fmt.Println(*pointr)

// ###### Arrays and Pointers
// Arrays are value types which mean a variable pass it to function, a copy of the entire array is made.
// We can use pointers to create a reference to an array, which allows to modify the original array.

  ar := [3]int{1,2,3}
  pttr := &ar
  fmt.Println(*pttr)

  // did not changed
  changeArray1(ar)
  fmt.Println(ar)

  // this change the actual values
  changeArray2(&ar)
  fmt.Println(ar)


  // ###### Functions and Pointers
  // Functions are one of the most common use cases for pointers.
  // We can use pointers to pass arguments to functions
  // Which allows to modify the function parameters
  //
  
  a := 1
  b := 2
  fmt.Println(a,b)
  swap(&a,&b)
  fmt.Println(a,b)
  
}

func addOne(x *int){
*x++;
}

func changeArray1(a [3]int){
  a[0] = 10
}

func changeArray2(a *[3]int){
  a[0] = 10
}

func swap(x *int, y *int){
  //  *x,*y = *y,*x
  temp := *x
  *x = *y
  *y = temp
}

