// Problem: Group Anagrams LeetCode #49
// Problem Statement
// Given an array of strings, group the anagrams together.
// You can return the answer in any order.

// Example:
// Input:
// strs = ["eat", "tea", "tan", "ate", "nat", "bat"]

// Output:
// [
//   ["eat","tea","ate"],
//   ["tan","nat"],
//   ["bat"]
// ]

// What is an Anagram?
// An anagram is a word formed by rearranging the letters of another word.

// For example:
// "eat", "tea", and "ate" → anagrams
// "tan" and "nat" → anagrams
// "bat" → no match

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(getAnagrams(s))
}

func getAnagrams(s []string) [][]string {
	hash := make(map[string][]string)

	for _, val := range s {
		sortedKey := sortWordAlphabetically(val)
		hash[sortedKey] = append(hash[sortedKey], val)
	}

	var result [][]string

	for _, val := range hash {
		result = append(result, val)
	}

	return result
}

func sortWordAlphabetically(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]

	})
	return string(runes)
}
