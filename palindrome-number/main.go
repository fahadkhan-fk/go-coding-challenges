package main

import (
	"fmt"
)

func main() {
	var n int = 101
	fmt.Println("is Palindrome: ", isPalindrome(n))
}

func isPalindrome(n int) bool {
	if n < 0 || (n%10 == 0 && n != 0) {
		return false
	}

	reversed := 0

	for n > reversed {
		reversed = reversed*10 + n%10
		n = n / 10
	}

	return reversed == n || n == reversed/10
}

// Time	Complexity O(log₁₀(n))
// Space Complexity	O(1)
