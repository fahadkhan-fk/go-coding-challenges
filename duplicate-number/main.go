// Problem:
// Given an unsorted array of integers, there is exactly one integer that appears twice. Your task is to find the integer that appears twice. The array can contain both positive and negative integers.

// Input:
// An unsorted array of integers, e.g., [3, 5, 2, 3, 1, 6].

// Output:
// The integer that appears twice, e.g., 3.

package main

import "fmt"

func main() {
	n := []int{3, 5, 2, 1, 3}
	hash := make(map[int]int)
	for _, v := range n {
		hash[v]++
		if hash[v] == 2 { // Found duplicate, print & exit
			fmt.Println("The integer that appears twice:", v)
			return
		}
	}
}

// Time Complexity: O(n)
// Space Complexity: O(n)
