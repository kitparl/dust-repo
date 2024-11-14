package main

import (
	"fmt"
	"time"
 	"sync"
	"net/http"
)
var wg sync.WaitGroup     // Pointers

// Concurrency
//concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or in partial order, without affecting the outcome
// Real example: Suppose i am eating a food and also i am playing games like bgmi. I am eating slice and till then i keep my phonne side and also vice versa

// Parallelism
// Parallelism is the ability of a program to run multiple tasks without interfering with each other.

// Goroutines
// Goroutins is the way to achive the parrallelism
/*
Threads: Managed by OS, fixed stack - 1MB
Goroutines: Managed by Go runtime, Fixed stack - 2KB
			
Note: Do not communicate by sharing memory; instead, share memory by communicating
*/
func main(){

	// go greeter("Hello")
	// greeter("Golang")
	
	websitelist := []string{
		"https://pranshu.tech",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

}

func greeter(s string) {

	for i := 0; i < 6; i++ {
		time.Sleep(3* time.Millisecond)
    fmt.Println(s)
  }

}

func getStatusCode(endpoint string){
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops error")
	}

	fmt.Println("%d status code for %s", res.StatusCode, endpoint)
}


