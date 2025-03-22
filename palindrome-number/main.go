// Problem: Palindrome Number (LeetCode #9)
// Problem Statement
// Given an integer x, return true if x is a palindrome, and false otherwise.

// A palindrome is a number (or string) that reads the same forward and backward.

package main

import (
	"fmt"
)

func main() {
	var x int = 101
	fmt.Println("is Palindrome: ", isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0

	for x > reversed {
		reversed = reversed*10 + x%10
		x = x / 10
	}

	return reversed == x || x == reversed/10
}

// Time	Complexity O(log₁₀(n))
// Space Complexity	O(1)
