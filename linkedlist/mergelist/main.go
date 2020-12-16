package main

import "fmt"

func main() {
	l1 := &linkedList{}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}
	l1.appendList(node1)
	l1.appendList(node2)
	l1.appendList(node3)

	l2 := &linkedList{}
	node4 := &ListNode{Val: 1}
	node5 := &ListNode{Val: 3}
	node6 := &ListNode{Val: 4}
	l2.appendList(node4)
	l2.appendList(node5)
	l2.appendList(node6)
	l1.printListValue()
	l2.printListValue()

	var nl *ListNode
	nl = mergeTwoLists(l1.head, l2.head)
	fmt.Println(nl)
}

type linkedList struct {
	head   *ListNode
	tail   *ListNode
	length int
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
func (l linkedList) printListValue() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf(" %v ", toPrint.Val)
		toPrint = toPrint.Next
		l.length--
	}
	fmt.Printf("\n")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList *ListNode
	pos := &ListNode{Val: -1, Next: nil}
	newList = pos

	for l1 != nil || l2 != nil {
		if l1.Val <= l2.Val {
			pos.Next = l1
			l1 = l1.Next
		} else {
			pos.Next = l2
			l2 = l2.Next
		}
		pos = pos.Next
	}

	// if some node is nil already, you can connect the entire of another node
	if l1 != nil {
		pos.Next = l1
	}
	if l2 != nil {
		pos.Next = l2
	}

	return newList.Next
}
