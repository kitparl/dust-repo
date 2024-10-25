package main


import (
    "errors"
    "fmt"
)

  //#######   Basic Error Handling
  // In Go errors are values
  type error interface {
    Error() string
  }

  func doSomething() error {
    return errors.New("Some error")
  }



  type MyError struct {
    Code int
    message string
  }
  func (e *MyError) Error() string {
    return fmt.Sprintf("%d - %s", e.Code, e.message) 
  }

  func doSomething2() error {
    return &MyError{Code: 404, message: "Not found"}
  }



// In Go, Errors are handled explicitly, typically using values of the built-in error type.
func main() {

  err := doSomething()
  if err != nil {
    fmt.Println(err)
  }


  //#######   Custom errors
  // Creating custom errors can be helpful if we wnat more structured information about the error.
  err2 := doSomething2()
  if err2 != nil {
    fmt.Println(err2)
  }


  
}
