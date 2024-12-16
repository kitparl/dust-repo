package main

import (
	"fmt"
)

func main() {
	str1 := "ram"
	fmt.Println("Address of variable str1:", &str1)

	str1 = "sham"
	fmt.Println("Address of variable str1 after reassignment:", &str1)


str3 := "hello"
str4 := "hello"
// Different variables have different addresses
fmt.Printf("%p\n", &str3)  // Different address
fmt.Printf("%p\n", &str4)  // Different address
}
