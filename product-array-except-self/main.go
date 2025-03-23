// Problem: Product of Array Except Self – LeetCode #238
// Problem Statement
// Given an integer array nums, return an array output such that:

// output[i] = product of all elements in nums **except nums[i]**

// Example:
// Input: nums = [1, 2, 3, 4]
// Output: [24, 12, 8, 6]

// Explanation:
// output[0] = 2 × 3 × 4 = 24
// output[1] = 1 × 3 × 4 = 12
// output[2] = 1 × 2 × 4 = 8
// output[3] = 1 × 2 × 3 = 6

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("Product of array except self: -> ", getProduct(nums))
}

func getProduct(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	prefix := 1
	suffix := 1

	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

// Time Complexity O(n)
// Space Complexity O(1)
