package main

import "fmt"

func main() {
	
	number := 5
	fmt.Println("Value of number in main is", number)
	fmt.Println()
	fmt.Println("Address of number in main is", &number)
	fmt.Println()
	case1(&number)
	fmt.Println("After case1, the value of number is", number)
}


func case1(number *int){ // here copy of the variable that is 5 is passed /// that is why this number holds a different address /// number variable is same but still different because of pointer (different address)
	// number = number + 10  ---> i cannot pointer + 10 ---> dereference it 
	
	fmt.Println("In case 1, value of number is", &number)
	*number = *number + 10

		fmt.Println("In case 1, address of number is", number)
	fmt.Println("In case 1, address of number is", *number)
	fmt.Println("In case 1, value of number is", &number)
}


// passing the value in arguments always a create a copy of the value of the vairable that is passed
// we can see the two differnt address. simply tells not the same
