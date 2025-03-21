// Problem Statement
// Given an integer array nums[], find the contiguous subarray (containing at least one number) which has the largest sum, and return that sum.

// This is a classic dynamic programming problem and is commonly solved using Kadane’s Algorithm.

// A subarray is a sequence of consecutive elements from the array.
// Example: For nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4], possible subarrays include:
// [1, -3, 4], [4, -1, 2, 1], [-2, 1], etc.

// Find the maximum possible sum of any contiguous subarray in nums.

// Example 1

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation:

// The subarray [4, -1, 2, 1] has the maximum sum: 4 + (-1) + 2 + 1 = 6

// Example 2

// Input: nums = [1]
// Output: 1
// The only subarray is [1].

package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))

}

func maxSubArray(arr []int) int {
	currentSum := arr[0]
	max := arr[0]
	for _, val := range arr[1:] {
		currentSum = maxx(val, currentSum+val)
		if currentSum > max {
			max = currentSum
		}
	}
	return max
}

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
