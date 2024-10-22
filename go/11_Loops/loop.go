
package main

import "fmt"

func main(){
/*
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
*/

for i:=0; i<5; i++{
  fmt.Println(i)
}


 for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }


for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }


// The Range Keyword

// The range keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value.

/*
for index, value := range array|slice|map {
   // code to be executed for each iteration
}
*/

fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
    fmt.Println(idx, val)
  }
// To only show the value or the index, you can omit the other output using an underscore (_).
  for _, val := range fruits {
     fmt.Println("val", val)
  }
}
