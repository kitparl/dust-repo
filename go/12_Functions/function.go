package main

import "fmt"

// creating a function
//1 use func keyword
//2 specify the function name
//3 add code declared in the function

func test(){
fmt.Println("Hi")
}
func main(){
// ### call the function test()
test()


// ### Function Parameters and Arguments

add(1,4) // function Arguments

// ### Catch return value
d := dummy()
println(d)

// ### catch multiple return value
a,b := multiReturnVal(3,"Pranshu")
fmt.Println(a,b)

//Note: If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.
_, d := multiReturnVal(5, "Hello")

}

func add(a int, b int){ // parameters and types
fmt.Println(a+b)
}

func dummy() int{
  // Function return values
  return 3;
}

func multiReturnVal(x int, y string) (int, string){
  return x, y
}
