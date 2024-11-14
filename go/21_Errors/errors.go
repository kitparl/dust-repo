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



  // #######    Custom Errors

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

  //####### 1. Basic error handling
  err := doSomething()
  if err != nil {
    fmt.Println(err)
  }


  //#######   2. Custom errors
  // Creating custom errors can be helpful if we wnat more structured information about the error.
  err2 := doSomething2()
  if err2 != nil {
    fmt.Println(err2)
  }

  //######## 3. Wrapping and Unwrapping errors
  // Wrapping errors allows us to wrap another error inside a new error, and then handle the wrapped error.
   
//   func doSomething() error {
//     err := errors.New("initial error")
//     return fmt.Errorf("failed to do something: %w", err)
// }

//  to check for the original error, use errors.Is and errors.As

// if errors.Is(err, initialError) {
//     fmt.Println("Initial error occurred")
// }
//
// var myErr *MyError
// if errors.As(err, &myErr) {
//     fmt.Println("Custom error with code:", myErr.Code)
// }
//

}
