// Problem statement: "Find the Missing Number"
// You are given an array of n distinct numbers where the numbers range from 0 to n.
// One number is missing, and your task is to find and return that missing number.

// Example:
// Input: nums = [3, 0, 1]
// Output: 2
// The numbers should be [0, 1, 2, 3], but 2 is missing.
// Expected output: 2.

package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}
	number := missingNumber(nums)
	fmt.Println("The missing number is: ", number)

}

func missingNumber(nums []int) int {
	n := len(nums)
	// mathematical formulae
	formulaSum := (n * (n + 1)) / 2
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return formulaSum - sum
}

// Time complexity O(n)
// Space complexity O(1)
