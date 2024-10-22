package main
import "fmt"

func main(){
  // Nested Loops

  adj := [2]string{"big", "small"}
  fruits := [3]string{"apple", "orange", "banana"}


  for i := 0; i<len(adj); i++ {
    for j := 0; j < len(fruits); j++ {
      fmt.Println(adj[i], fruits[j])
    }
  }
}
