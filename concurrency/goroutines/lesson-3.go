package main

import (
	"fmt"
	"sync"
)

// wg.Add(1) => one goroutine to wait for
// defer wg.Done() => mark it finished
// wg.Wait() => block until all done

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("Numbers: i = ", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	wg.Wait()

	for j := 1; j <= 5; j++ {
		fmt.Println("Numbers: j = ", j)
	}
}
