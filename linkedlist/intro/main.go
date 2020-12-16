package main

import "fmt"

type node struct {
	prev  *node
	next  *node
	value interface{}
}

type linkedList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func (l *linkedList) prepend(n *ListNode) {
	second := l.head
	l.head = n
	l.head.Next = second
	l.length++
}

func (l *linkedList) appendList(n *ListNode) {
	if l.head == nil {
		l.head = n
		l.length++
		return
	}
	currNode := l.head
	for currNode != nil {
		if currNode.Next == nil {
			currNode.Next = n
			break
		}
		currNode = currNode.Next
	}
	l.length++
}

// func (l *linkedList) deleteWithValue(value interface{}) {
// 	previousToDelete := l.head
// 	for previousToDelete.next.value != value {
// 		previousToDelete = previousToDelete.next
// 	}
// 	previousToDelete.next = previousToDelete.next.next
// 	l.length--
// }

func (l linkedList) printListValue() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf(" %v ", toPrint.Val)
		toPrint = toPrint.Next
		l.length--
	}
	fmt.Printf("\n")
}

// func (l *linkedList) reverse() {
// 	currNode := l.head
// 	var prev *node
// 	for currNode != nil {
// 		currNode, prev, currNode.prev = currNode.prev, currNode, prev
// 	}
// 	l.head = prev
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	currNode := head
	var prev *ListNode
	var next *ListNode

	for currNode != nil {
		next = currNode.Next
		currNode.Next = prev
		prev = currNode
		currNode = next
	}
	return prev
}

func main() {
	mylist := &linkedList{}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	mylist.appendList(node1)
	mylist.appendList(node2)
	mylist.appendList(node3)
	mylist.appendList(node4)
	mylist.appendList(node5)
	mylist.printListValue()
}
