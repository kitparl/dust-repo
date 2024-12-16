package main

import "fmt"

func main() {
	
	number := 5

	fmt.Println("Value of number in main is", number)
	case1(number)
	fmt.Println("Address of number in main is", &number)
	fmt.Println("After case1, the value of number is", number)
}


func case1(number int){ // here copy of the variable that is 5 is passed
	number = number + 10
	fmt.Println("In case 1, address of number is", number)
	fmt.Println("In case 1, value of number is", &number)
}


// passing the value in arguments always a create a copy of the value of the vairable that is passed
// we can see the two differnt address. simply tells not the same
