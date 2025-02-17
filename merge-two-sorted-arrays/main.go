// Problem Statement
// Given two sorted arrays, and your task is to merge them into one sorted array
// while maintaining the order.
// Example:
// arr1 = [1, 3, 5, 7]
// arr2 = [2, 4, 6, 8]
// Expected Output:
// [1, 2, 3, 4, 5, 6, 7, 8]

package main

import "fmt"

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8}
	finalArr := mergeSortedArrays(arr1, arr2)
	fmt.Println(finalArr)
}

// Time complexity O(n + m) ; Space Complexity O(n + m)
func mergeSortedArrays(arr1, arr2 []int) []int {
	var arr3 []int
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		} else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}

	if i < len(arr1) {
		arr3 = append(arr3, arr1[i])
		i++
	}

	if j < len(arr2) {
		arr3 = append(arr3, arr2[j])
		j++
	}
	return arr3
}
