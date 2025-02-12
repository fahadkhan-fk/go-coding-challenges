// Palindrome Check "racecar" "madam"
// Write a function that checks if a given string is a palindrome.
// func isPalindrome(s string) bool
package main

import (
	"fmt"
)

func main() {
	s := "abbbba"
	fmt.Println("is palindrome:", isPalindrome(s))
}

// NOTES: O(n) time complexity,  O(1) space complexity, Iterative approach
func isPalindrome(s string) bool {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] { // directly use index math instead of `j`
			return false
		}
	}
	return true
}
