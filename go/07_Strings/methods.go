package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "tum"
	str2 := "hum"

	// Compare	
	// if a < b = -1 , a > b = -1, a = b = 1
	res := strings.Compare(str1, str2)

	fmt.Println("compare result --> ", res)


	// Contains
	
	result := strings.Contains(str1, "tu")
  fmt.Println("contains result --> ", result)


	// Replace
	
	str2 = strings.Replace(str1, "t", "h", 1)
	fmt.Println("replace result --> ", str2)

	// ToUpper()
	
	newStr1 := strings.ToUpper(str1)
	fmt.Println("toUpper result --> ",newStr1)


	// ToLower()
	newStr1 = strings.ToLower(newStr1)
	fmt.Println(newStr1)
	fmt.Println("toLower result --> ",newStr1)


	// Split()
	phrase := "Pranshu is my name"
	split := strings.Split(phrase, " ")
  fmt.Println("split result --> ",split) // slice


	// Compare strings
	compareRes := str1 == str2

	fmt.Println("compare result --> ", compareRes)


	// Create strings from a slice (Join)
	words := []string{"Pranshu", "Singh", "Bisht"}
	msg := strings.Join(words, " ")
	fmt.Println(msg)

	// Iteration over the strings
	var name string = "pranshu"

	for i:=0; i<len(name); i++ {
		fmt.Print(string(name[i]))
	}


	for i, e := range name {
    fmt.Println(i, e)
  }


	fmt.Println(string(97))


}
