package main

import (
	"crypto/rand"         // For cryptographic randomness
	"fmt"
	"math/big"            // For handling big integers
	mathrand "math/rand"  // Aliased "math/rand" to avoid conflict
	"time"
)

func main() {

	// Random number using math/rand
	mathrand.Seed(time.Now().UnixNano())
	fmt.Println("random number:", mathrand.Intn(5)+1)

	// Random number using crypto/rand
	myRandomNumber, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}
	fmt.Println("cryptographic random number:", myRandomNumber)
}
