// Problem Statement: Is valid parentheses
// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//	Example:
//
// Input: "{[()]}"
// Output: true ✅

// Input: "{[(])}"
// Output: false ❌

// func isValid(s string) bool
package main

import (
	"fmt"
	"strings"
)

func main() {
	parentheses := "{}[]"
	fmt.Println(isValid(parentheses))
}

func isValid(s string) bool {
	lens := len(strings.TrimSpace(s))
	if lens == 0 {
		return true
	}

	if lens%2 != 0 {
		return false
	}

	stack := []rune{}

	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '{' || char == '(' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if bracketPairs[char] != top {
				return false
			}
		}
	}
	return len(stack) == 0
}
