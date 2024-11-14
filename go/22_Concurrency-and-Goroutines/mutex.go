package main

import (
	"fmt"
  "sync"
	"net/http"
)

// In Go (Golang), a Mutex is a synchronization primitive that provides a way to safely manage access to shared resources across multiple goroutines.
// Mutex stands for Mutual Exclusion. It allows only one goroutine to enter the critical section, ensuring exclusive access to the resource.
// It is part of the sync package, which provides basic synchronization primitives in Go.

/*
A mutex has two primary methods:

- Lock(): Acquire a lock, preventing other goroutines from accessing the shared resource.
- Unlock(): Release the lock, allowing other goroutines to access the shared resource.
*/


var wg sync.WaitGroup     // Pointers
var mut sync.Mutex        // Pointers
var signals = []string{"test"}

func main(){

	
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

func getStatusCode(endpoint string){
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops error")
	}else{
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Println("%d status code for %s", res.StatusCode, endpoint)
	}

}


