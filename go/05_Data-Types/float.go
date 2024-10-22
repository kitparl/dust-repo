package main

import "fmt"


func main(){
   
/*
  # float32
  ## size
  32 bits
  ## Range
 -3.4e+38 to 3.4e+38.

  # float64
  ## size
  64 bits
  ## Range
 -1.7e+308 to +1.7e+308.

 NOTE: The default type for float is float64
 */
 
 var x float32 = 123.78
 var y float32 = 3.4e+38
 fmt.Printf("Type: %T, value: %v\n", x, x)
 fmt.Printf("Type: %T, value: %v", y, y)


 var a float64 = 1.7e+308
 fmt.Printf("Type: %T, value: %v", a, a)
}
