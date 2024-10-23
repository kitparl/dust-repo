package main

import "fmt"

func main() {
  /* Slice or Arrays */

  //  I used range to sum the numbers in a slice. Arrays work like this too.

  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)
  
  // range on arrays and slices provides both the index and value for each entry.
  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i) // index: 1
    }
  }

  /* map */
  myMap := map[string]int{"a": 1, "b": 2}
  for k, v := range myMap {
    fmt.Println(k, v)
  }

  for k := range myMap {
    fmt.Println(k)
  }


  /* strings */
  // range on strings iterates over Unicode code points
  for i, c := range "go" {
    fmt.Println(i, c)
  }


}
