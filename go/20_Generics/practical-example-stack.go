package main
 
import "fmt"

type Stack[T any] struct {
  elements []T
}

func (s *Stack[T]) Push(value T) {
  s.elements = append(s.elements, value)
}

func (s *Stack[T]) Pop() T {
  if len(s.elements) == 0 {
    var zero T
    return zero
  }

  last := s.elements[len(s.elements)-1]
  s.elements = s.elements[:len(s.elements)-1]
  return last
}

func main() {

  intStack := Stack[int]{}
  intStack.Push(1)
  intStack.Push(5)
  fmt.Println(intStack)
  fmt.Println(intStack.Pop())

}


// Generics are powerful in Go, providing the flexibility to write generic algorithms and structures that are type-safe and efficient
