package main

func main() {

}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	currNode := root
	var currChild *Node
	var next *Node

	// top-down solution
	for currNode != nil {
		// merge
		currChild = currNode.Child
		next = currNode.Next

		if currChild != nil {
			currNode.Child = nil
			currNode.Next = currChild
			currChild.Prev = currNode

			// travel to the last child
			if currChild.Next != nil {
				for currChild.Next != nil {
					currChild = currChild.Next
				}
			}
			if next == nil {
				if currChild.Child != nil {
					childOfChild := currChild.Child
					currChild.Next = childOfChild
					childOfChild.Prev = currChild
					currChild.Child = nil
				}
			} else {
				currChild.Next = next
				next.Prev = currChild
			}
		} else {
			currNode = currNode.Next
		}
	}

	return root
}
