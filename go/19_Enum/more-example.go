package main

import "fmt"

/*
const (
  c0 = iota
  c1 = iota
  c2 = iota

)
*/


// can avoid writing successive iota in front of every constant.

const (
  c0 = iota
  c1
  c2
  _      // _ for skip
  c4
)

func main() {
  fmt.Println(c0, c1, c2, c4) // 0, 1, 2
}
