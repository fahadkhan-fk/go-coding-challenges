// Problem Statement
// Given two sorted link list, and your task is to merge them into one sorted list
// while maintaining the order.
// Example:
// arr1 = [1, 3, 5, 7]
// arr2 = [2, 4, 6, 8]
// Expected Output:
// [1, 2, 3, 4, 5, 6, 7, 8]

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{10, 11, 12, 13}

	list1 := createLinkedList(arr1)
	list2 := createLinkedList(arr2)

	sortedList := sortLinkList(list1, list2)

	printLinkedList(sortedList)

}

func printLinkedList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val)
		if list.Next != nil {
			fmt.Print(" -> ")
		}
		list = list.Next
	}
	fmt.Println()
}

func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func sortLinkList(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	current := l3
	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return l3.Next
}
