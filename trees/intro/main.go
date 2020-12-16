package main

import "fmt"

func main() {
	root := &Node{9, nil, nil}
	n1 := &Node{4, nil, nil}
	n2 := &Node{6, nil, nil}
	n3 := &Node{12, nil, nil}
	n4 := &Node{170, nil, nil}
	n5 := &Node{15, nil, nil}
	n6 := &Node{1, nil, nil}

	root.Insert(n1)
	root.Insert(n2)
	root.Insert(n3)
	root.Insert(n4)
	root.Insert(n5)
	root.Insert(n6)
	fmt.Println(root.lookup(4))
	traverse(root)

}

type Tree struct {
	root *Node
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(node *Node) {
	for true {
		if node.Val < n.Val {
			if n.Left == nil {
				n.Left = node
				break
			}
			n = n.Left
		} else {
			if n.Right == nil {
				n.Right = node
				break
			}
			n = n.Right
		}
	}
	return
}

func traverse(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Val)
		traverse(n.Left)
		traverse(n.Right)
	}
}

func (n *Node) lookup(val int) *Node {
	for n != nil {
		if val < n.Val {
			// go left
			n = n.Left
		} else if val > n.Val {
			// go right
			n = n.Right
		} else {
			return n
		}
	}

	return nil
}

// func (n *Node) remove(val int) *Node {
// 	for n != nil {
// 		if val < n.Val {
// 			// go left
// 			n = n.Left
// 		} else if val > n.Val {
// 			// go right
// 			n = n.Right
// 		} else if val == n.Val && n.Left == nil && n.Right == nil {
// 			// Leaf case
// 			elm := n
// 			n = nil
// 			return elm
// 		} else if val == n.Val && n.Right == nil
// 	}

// 	return nil
// }
