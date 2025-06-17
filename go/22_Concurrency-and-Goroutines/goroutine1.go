package main

import (
 "fmt"
 "time"
) 


func printMessage(msg string) {
 fmt.Println(msg)
} 

func main() { 
// keyword go added in front of calling method will run the function in a sparate excution flow.
 go printMessage("Hello, Goroutine!")
// Internally Above will be part of 'run queue' which is managed by Go Scheduler
 fmt.Println("Main function message")
 time.Sleep(time.Second) // waiting for goroutine to finish. 
}
