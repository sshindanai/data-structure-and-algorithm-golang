package main

func main() {

}

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func reverseBetween(head *ListNode, m int, n int) *ListNode {
// 	if m == n {
// 		return head
// 	}
// 	// currPosition: Tell us what we need to do on each node
// 	currPosition := 1

// 	currNode := head

// 	// start: Set to be the node before reverse start
// 	start := head

// 	for currPosition < m {
// 		start = currNode
// 		currNode = currNode.Next
// 		currPosition++
// 	}

// 	// start to reverse linked list
// 	// tail: Equal to the first value in the start of reversal
// 	tail := currNode
// 	var newList *ListNode
// 	var next *ListNode
// 	for currPosition >= m && currPosition <= n {
// 		next = currNode.Next
// 		currNode.Next = newList
// 		newList = currNode
// 		currNode = next
// 		currPosition++
// 	}

// 	start.Next = newList
// 	tail.Next = currNode

// 	if m > 1 {
// 		return head
// 	}
// 	return newList
// }

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	prevOfPrev := 0
	prev := 1
	curr := 0

	for i := 0; i < n-1; i++ {
		curr = prevOfPrev + prev
		prevOfPrev = prev
		prev = curr
	}

	return curr
}
