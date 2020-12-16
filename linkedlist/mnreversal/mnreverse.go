package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	currPos := 1
	currNode := head
	// start stores before reverse nodes
	start := head

	for currPos < m {
		start = currNode
		currNode = currNode.Next
		currPos++
	}

	var newList *ListNode
	var next *ListNode
	tail := currNode

	for currPos >= m && currPos <= n {
		next = currNode.Next
		currNode.Next = newList
		newList = currNode
		currNode = next
		currPos++
	}

	// connect
	start.Next = newList
	tail.Next = currNode

	if m > 1 {
		return head
	}
	// m == 1
	return newList
}
