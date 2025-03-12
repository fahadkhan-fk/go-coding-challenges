// Binary Search Problem
// You need to implement binary search on a sorted array to efficiently find the index of a target number.
// Example 1:
// Input: nums = [2, 5, 8, 12, 16, 23, 38, 45, 56, 72, 81, 90, 101, 110, 125], target = 12
// Output: 2

package main

import "fmt"

func main() {
	target := 12
	nums := []int{2, 5, 8, 12, 16, 23, 38, 45, 56, 72, 81, 90, 101, 110, 125}
	index := binarySearch(target, nums)
	fmt.Println(index)

}

func binarySearch(target int, nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid // found the target
		} else if nums[mid] < target {
			left = mid + 1 // move right
		} else {
			right = mid - 1 // move left
		}
	}
	return -1
}

// O(log n) time complexity
// More memory-efficient (O(1) space complexity)
