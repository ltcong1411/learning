package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new channel
	ch := make(chan string)

	// Start a goroutine
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from Goroutine"
	}()

	// Main goroutine waits to receive data
	message := <-ch
	fmt.Println(message)
}
