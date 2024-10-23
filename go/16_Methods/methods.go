package main

import "fmt"

// Methods are similar to Go functions with one difference, i.d the method contains a reciever argument in it.
// The reciever is the special parameter that acts as the context for the method.

type Rectangle struct {
  width, height int
}

// define a method
func (r Rectangle) Area() int {
  return r.width * r.height
}

// define the method that modifies the reciever
func (r *Rectangle) setWidthHeight(factor int) {
  r.width = factor
  r.height = factor
}

func (r *Rectangle) setWidth(factor int) {
  r.width = factor
}

func (r *Rectangle) setHeight(factor int) {
  r.height = factor
}

func main() {
  rect := Rectangle{width: 4, height: 5}
  fmt.Println(rect.Area())

  rect.setWidth(6)
  fmt.Println(rect.Area())

  rect.setHeight(2)
  fmt.Println(rect.Area())

}



// Methods enable more structured and organized code by encapsulating behaviour within the type.
// it's common to use structs as receivers, a receiver can be any type, including: Structs, Pointer to structs, Any user defined type, Non structs like slice, arrays etc.
