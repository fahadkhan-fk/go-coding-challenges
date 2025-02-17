// Given an array of integers, return indices of the two numbers
// such that they add up to a specific target.
// func twoSum(nums []int, target int) []int
// nums = [2, 7, 11, 15]
// target = 9
// The numbers 2 and 7 add up to 9. Their indices are [0, 1]. Expected Output: [0, 1]
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	indices := twoSum(nums, 22)
	fmt.Println("Indices are: ", indices)
}

// Below solution works but Time Complexity: O(n²) due to nested loops.
// func twoSum(nums []int, target int) []int {
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ { // Start from i+1 to avoid self-addition
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// O(n) Time Complexity, O(n) Space Complexity, Efficient for Large Arrays
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // map to store [number]index

	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i} // found the solution
		}
		numMap[num] = i // store the current number and its index
	}
	return []int{}
}
