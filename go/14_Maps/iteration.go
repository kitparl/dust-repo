package main
import ("fmt")

func main() {
  //##### Iterate Over Maps
// we can use range to iterate over maps.


  example := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  for k, v := range example {
    fmt.Println(k, v)
  }

// ###### Iterate over Maps on specfic order

  var b []string             // defining the order
  b = append(b, "one", "two", "three", "four")

  for k, v := range example {        // loop with no order
    fmt.Printf("%v : %v, ", k, v)
  }

  fmt.Println()

  for _, element := range b {  // loop with the defined order
    fmt.Printf("%v : %v, ", element, example[element])
  }
}
