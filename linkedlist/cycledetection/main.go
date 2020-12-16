package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	currNode := head
	has := make(map[*ListNode]int)
	for currNode != nil {
		has[currNode]++
		if has[currNode] > 1 {
			return currNode
		}
		currNode = currNode.Next
	}
	return nil
}
