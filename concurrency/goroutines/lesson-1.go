package main

import (
	"fmt"
	"time"
)

// printNumbers() runs in a separate goroutine
// main() keeps running in parallel with it
// both outputs may interleave

// When main() exits, the whole program exits, even if goroutines are still running.
// That is why goroutines often need coordination and we added that here below through time.Sleep(200 * time.Millisecond)

func hello() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number i:", i)
		time.Sleep(200 * time.Millisecond)
	}
}
func main() {
	go hello()

	for j := 1; j <= 5; j++ {
		fmt.Println("Number j:", j)
		time.Sleep(200 * time.Millisecond)
	}
}
