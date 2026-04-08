package main

import "fmt"

// This may print nothing.
// Why? Because main() can finish before the goroutine gets CPU time.
// That leads us to the next important tool: WaitGroup.

func hello() {
	fmt.Println("Hello from goroutines")
}

func main() {
	go hello()
}
