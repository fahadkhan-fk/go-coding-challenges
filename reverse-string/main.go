// Reverse a String
// Implement a function that reverses a string.
// func reverseString(s string) string
package main

import "fmt"

func main() {
	word := "golang"
	reversedString := reverseString(word)
	fmt.Println(reversedString)
}

// Below solution works but the issue is optimization as rs += returns new string each time; This results in an O(n²) complexity instead of the optimal O(n).
// func reverseString(s string) string {
// 	var rs string
// 	length := len(s) - 1

// 	for i := length; i >= 0; i-- {
// 		rs += string(s[i])
// 	}
// 	return rs
// }

// Better approach; O(n) Time Complexity;
func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
