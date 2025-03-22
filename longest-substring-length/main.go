// Problem Statement
// Given a string, and you need to find the length of the longest substring
// that contains no repeating characters.
// Example:
// s = "abcabcbb"
// Substrings Without Repeating Characters: "abc" -> Length 3

package main

import "fmt"

func main() {
	len := lengthOfLongestSubstring("abcabcbb")
	fmt.Println("longest substring length: ", len)
}

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	maxLength := 0
	start := 0 // left boundary of the sliding window
	for end, value := range s {
		// if character already exists in the map AND is within the window, move `start`
		if index, exists := charMap[value]; exists && index >= start {
			start = index + 1
		}
		// store/update the last seen position of the character
		charMap[value] = end
		// update maxLength
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

// Alternative approach
// func isUnique(s string) int {
// 	hash := make(map[rune]bool)
// 	left := 0
// 	max := 0

// 	for idx, r := range s {
// 		if hash[r] == true {
// 			left++
// 			continue
// 		}

// 		if !hash[r] {
// 			hash[r] = true
// 			if max < (idx+1)-left {
// 				max = (idx + 1) - left
// 			}
// 		}
// 	}
// 	return max
// }
