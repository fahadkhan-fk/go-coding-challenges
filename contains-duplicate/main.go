// Problem Statement
// You are given an array of integers nums[]. Your task is to determine if any value appears at least twice in the array.

// Return:
// true → if any number appears more than once.
// false → if all numbers are unique.

// Example 1

// Input: nums = [1,2,3,1]
// Output: true
// 1 appears twice, so return true.

// Example 2

// Input: nums = [1,2,3,4]
// Output: false
// All numbers are unique, so return false

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(duplicateExists(nums))
}

func duplicateExists(n []int) bool {
	uniqueMap := make(map[int]struct{})
	for _, value := range n {
		if _, ok := uniqueMap[value]; ok {
			return true
		}
		uniqueMap[value] = struct{}{}
	}
	return false
}

// Time	O(n) – Each number is checked once
// Space	O(n) – In worst case, all numbers are unique and stored in the map
