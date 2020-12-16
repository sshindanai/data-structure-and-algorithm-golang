package solution

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
	for currNode != nil {
		if currNode.Child == nil {
			currNode = currNode.Next
		} else {
			currChild := currNode.Child
			// loop to the last
			for currChild.Next != nil {
				currChild = currChild.Next
			}

			currChild.Next = currNode.Next
			// if the parent list still has node left
			if currChild.Next != nil {
				// set [currentChild] <- [next parent node]
				currChild.Next.Prev = currChild
			}

			currNode.Next = currNode.Child
			currNode.Next.Prev = currNode
			currNode.Child = nil
		}
	}

	return root
}
