package main

import "fmt"

// A struct (short for structure) is used to create collection of members of different data types, into a single variable.
// While arrays are used to store multiple values of the same type in a single variable, in struct, we can declare multiple variables of different data types.
// A struct can be useful for grouping data together to create records.

// To declare a structure in Go, use the type and struct keywords:

type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {
  // Access Struct Members

  var person1 Person
  var person2 Person

  person1.name = "Golang"
  person1.age = 20
  person1.job = "Engineer"
  person1.salary = 100000


  person2.name = "Go"
  person2.age = 15
  person2.job = "Developer"
  person2.salary = 200000

  fmt.Println(person1.name)
  fmt.Println(person1.age)
  fmt.Println(person1.job)
  fmt.Println(person1.salary)

  fmt.Println(person2.name)
  fmt.Println(person2.age)
  fmt.Println(person2.job)
  fmt.Println(person2.salary)


  // Pass Struct as Function Arguments
  printPerson(person1)
  printPerson(person2)

}

func printPerson(p Person){
  fmt.Println(p.name)
  fmt.Println(p.age)
  fmt.Println(p.job)
  fmt.Println(p.salary)
}
