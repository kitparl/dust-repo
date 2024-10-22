
package main

import "fmt"

func main(){

  day := 4


  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
    fmt.Println("Invalid Day")
  }

  // The default keyword specifies some code to run if there is no case match


  // multicase switch case
  switch day {
    case 1, 2, 3, 4, 5:
      fmt.Println("Weekday")
    case 6, 7:
      fmt.Println("Weekend")
    default:
      fmt.Println("Invalid Day")
  }
}
