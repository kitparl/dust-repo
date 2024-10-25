package main

import "fmt"


// In Go, struct embedding is a technique that allows to include struct within another struct.
// Effectively enableing "inheritance-like" behaviour in Go.


type Person struct {
  Name string
  Age int
}

type Employee struct {
  Person // Embedded struct
  Position string
  Salary int
}

func (p Person) Greet() {
  fmt.Printf("I am %s and I am %d year old. \n", p.Name, p.Age)
}

func main() {

  // create an Employee instance
  emp := Employee{
    Person: Person{
      Name: "Pranshu",
      Age: 24,
    },
    Position: "Software Engineer",
    Salary: 80000,
  }

  fmt.Println("Employee", emp)
  fmt.Println("Employee Name", emp.Name)
  fmt.Println("Age", emp.Age)
  fmt.Println("Position", emp.Position)

  emp.Greet()
  
}
