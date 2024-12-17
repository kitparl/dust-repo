package main

import (
	"fmt"
)

func main() {
	str1 := "pranshu"
	fmt.Println("Address of variable str1:", &str1)

	str1 = "pranshu"+"a"
	fmt.Println("Address of variable str1 after reassignment:", &str1)
}
