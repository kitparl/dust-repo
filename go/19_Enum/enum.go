package main

import "fmt"

// Define a new type for enum
type Day int

const (
  Sunday Day = iota // iota creates successive integer constatns starting from 0, so Sunday = 0, Monday = 1..... so on
  Monday
  Tuesday
  Wednesday
  Thursday
  Friday
  Saturday
)

func (d Day) String() string {
  return []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}

func main() {
  var today Day = Friday
  fmt.Println(today)


  if today == Friday {
    fmt.Println("It's friday")
  }
}
