package main

import "fmt"

func main(){

  // ## if statment
  if 20> 19 {
    fmt.Println("Correct")
  }


  x:=45
  y:=76

  if x < y {
    fmt.Println("y is greater")
  }


  // ## if-else statement
  if 20 > 19 {
    fmt.Println("Hi Golang")
  } else {
    fmt.Println("Bye Golang")
  }


  age := 10
  if(age != 18){
    fmt.Println("You are not 18")
  }else{ //  else brackets should be in same line  otherwise in a different line will raise an error
    fmt.Println("You are 18")
  }


  // ## if-else if-else

  time := 12

  if(time < 12){
    fmt.Println("Good morning")
  }else if(time < 18){
    fmt.Println("Good afternoon")
  }else{
    fmt.Println("Good evening")
  }

  // ## Nested if statement
  
  n := 5

  if(n>0){
    if(n<10){
      fmt.Println("N is less than 10")
    }else{
      fmt.Println("N is greater than 10")
    }
  }

}
