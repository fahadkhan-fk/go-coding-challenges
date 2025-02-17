// Problem Statement
// The factorial of a number n is the product of all positive integers from 1 to n.
// 5! = 5 × 4 × 3 × 2 × 1 = 120
package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println("factorial: ", fact)
}

// Time Complexity O(n) ; Space Complexity O(1)
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fact := 1
	for i := n; i > 0; i-- {
		fact = fact * i
	}
	return fact
}
