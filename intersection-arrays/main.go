// Problem: Intersection of Two Arrays (LeetCode #349)
// Problem Statement
// You are given two integer arrays, nums1 and nums2.

// Your task is to return an array of their intersection, meaning:

// Each element in the result must be unique.
// The result can be returned in any order.

// Example:1
// Input:
// nums1 = [1, 2, 2, 1]
// nums2 = [2, 2]
// Output: [2]

// Example:2
// Input:
// nums1 = [4,9,5]
// nums2 = [9,4,9,8,4]
// Output: [9, 4] or [4, 9] (any order) 'order does not matter'

package main

import "fmt"

func main() {
	arr1 := []int{4, 9, 5}
	arr2 := []int{9, 4, 9, 8, 4}

	arr3 := getIntersection(arr1, arr2)
	fmt.Println("Intersection array: ", arr3)

}

func getIntersection(arr1, arr2 []int) []int {
	seen := make(map[int]bool)
	arr := []int{}
	alreadyAdded := make(map[int]bool)

	for _, val := range arr1 {
		seen[val] = true
	}

	for _, val := range arr2 {
		if seen[val] && !alreadyAdded[val] {
			arr = append(arr, val)
			alreadyAdded[val] = true
		}
	}

	return arr

	// below approach works but contains one extra for loop

	// hash := make(map[int]bool)
	// hash2 := make(map[int]struct{})
	// uniqueArr := []int{}

	// i, j := 0, 0

	// for i < len(arr1) {
	// 	hash[arr1[i]] = true
	// 	i++
	// }

	// for j < len(arr2) {
	// 	if _, ok := hash[arr2[j]]; ok {
	// 		hash2[arr2[j]] = struct{}{}
	// 	}
	// 	j++
	// }

	// for key, _ := range hash2 {
	// 	uniqueArr = append(uniqueArr, key)
	// }

	// return uniqueArr
}

// Time Complexity O(n + m)
// Space Complexity O(n)
