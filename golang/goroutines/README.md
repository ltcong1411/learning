# Goroutines

### When would you use goroutines in Golang? Can you explain how channels work and how they help synchronize data between goroutines?

In Go, **goroutines** are lightweight threads used for concurrent execution of functions. You would use goroutines in situations where you want to perform multiple tasks simultaneously, such as:

1. **I/O operations**: When waiting for external input (e.g., file reading/writing, database queries, or network calls), goroutines allow other tasks to proceed without blocking.
   
2. **Parallel processing**: If you have CPU-bound tasks that can run independently, like image processing, data transformations, or calculations, you can run them concurrently to improve performance on multicore systems.

3. **Event-driven systems**: When building systems that handle multiple events or requests concurrently, like web servers, goroutines are ideal because they handle each request without blocking the main process.

### How Channels Work

**Channels** in Go are used to synchronize data between goroutines by allowing them to communicate in a safe and structured way. Channels can be thought of as pipes where one goroutine sends data and another receives it.

Here's an overview of how channels work:

- **Creating a Channel**: You create a channel using the `make` function:
  ```go
  ch := make(chan int)
  ```
  This creates a channel that can carry `int` data.

- **Sending Data**: A goroutine can send data to a channel using the `<-` operator:
  ```go
  ch <- 42
  ```
  This sends the value `42` into the channel `ch`.

- **Receiving Data**: A goroutine can receive data from a channel using the same `<-` operator:
  ```go
  value := <- ch
  ```
  This receives the value from the channel `ch` and stores it in the variable `value`.

### Synchronization with Channels

Channels help synchronize goroutines by blocking operations until both the sender and receiver are ready:

- **Sending** blocks the goroutine until another goroutine is ready to receive from the channel.
- **Receiving** blocks the goroutine until another goroutine sends data into the channel.

This built-in blocking behavior ensures that goroutines are naturally synchronized without needing explicit locking mechanisms like mutexes.

### Example

```go
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
```

In this example, the main goroutine waits until the second goroutine sends the message `"Hello from Goroutine"` into the channel after a 2-second delay. The synchronization happens naturally through the channel, avoiding the need for manual synchronization techniques.

### Buffered Channels

You can also create **buffered channels**, which allow sending and receiving without blocking, up to a specified limit:
```go
ch := make(chan int, 3) // Buffered channel with a capacity of 3
```
This allows up to 3 values to be sent to the channel without blocking the sender until the buffer is full.