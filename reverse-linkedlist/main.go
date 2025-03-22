// Problem: Reverse a Linked List (LeetCode #206)
// Problem Statement
// Given the head of a singly linked list, reverse the list, and return the new head.

// Input:  1 → 2 → 3 → 4 → 5
// Output: 5 → 4 → 3 → 2 → 1

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := createLinkedList(arr)
	printLinkedList(list)
	printLinkedList(reverseLinkedList(list))
}

func reverseLinkedList(list *ListNode) *ListNode {
	var prev *ListNode = nil
	current := list
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev

}

func createLinkedList(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head

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
