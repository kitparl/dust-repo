package main

import "fmt"

// The integer data type has two categories
// 1. Signed Integer : can store both +ve and -ve numbers
// 2. Unsigned Integer : can store only non -ve numbers


// The default type for integer is int. If you do not specify a type, the type will be int.

func main(){

  // Signed Integers

  var x int = 500
  var y int = -4500

  fmt.Printf("Type: %T, value: %v ", x, x)
  fmt.Printf("Type: %T, value: %v ", y, y)

  // Go has five keywords/types of signed integers
  // int, int8, int16, int32, int64

  /*
  # int
  ## size
 Depends on platform: 32 bits in 32 bit systems and 64 bit in 64 bit systems
  ## Range
-2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems


  # int8
  ## size
  8 bits/1 byte
  ## Range
  -128 to 127


  # int16
  ## size
  16 bits/2 byte
  ## Range
  -32768 to 32767


  # int32
  ## size
  32 bits/4 byte
  ## Range
 	-2147483648 to 2147483647


  # int64
  ## size
  64 bits/8 byte
  ## Range
	-9223372036854775808 to 9223372036854775807
  */

   
// Unsigned Integers
// Unsigned integers, declared with one of the uint keywords, can only store non-negative values

  var a uint = 500
  var b uint = 4500
  fmt.Printf("Type: %T, value: %v", a, a)
  fmt.Printf("Type: %T, value: %v", b, b)

// Go has five keywords/types of unsigned integers:

/*

  # uint
  ## size
  Depends on platform: 32 bits in 32 bit systems and 64 bit in 64 bit systems
  ## Range
	0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems*/

  # uint8 
  ## size
  8 bits/1 byte
  ## Range
	0 to 255

  # uint16 
  ## size
  16 bits/2 byte
  ## Range
	0 to 65535

  # uint32 
  ## size
  32 bits/4 byte
  ## Range
	0 to 4294967295


  # uint64 
  ## size
  64 bits/8 byte
  ## Range
	0 to 18446744073709551615
  */

  // Which Integer Type to Use?
  // The type of integer to choose, depends on the value the variable has to store.
}
