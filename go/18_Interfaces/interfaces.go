package main

import "fmt"

// an interface is a type that defines a set of method signatures but does not implement them.
// This concept is central to Go's approach to polymorphism.


// #### Define an interface named shape
type Shape interface {
  Area() float64
}

// Define struct for circle
type Circle struct {
  Radius float64
}

// Implement the Area method
func (c Circle) Area() float64 {
  return 3.14 * c.Radius * c.Radius
}

// Define a struct for Rectangle
type Rectangle struct {
  Width float64
  Height float64
}

// Implement the Area method for rectangle
func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func main() {
  // Create instances of circle and rectangle
  c := Circle{Radius: 5.0}
  r := Rectangle{Width: 3.0, Height: 4.0}

  // interface variable holding different types
  var s Shape

  s = c // Circles satifies Shape
  fmt.Println(s.Area())

  s = r // Rectangles satifies Shape
  fmt.Println(s.Area())
}
