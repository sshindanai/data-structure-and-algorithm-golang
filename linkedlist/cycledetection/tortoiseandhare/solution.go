package tortoiseandhare

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

	tortoise := head.Next
	hare := head.Next.Next

	for tortoise != hare {
		if hare == nil || hare.Next == nil {
			return nil
		}
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}

	tortoise = head

	for tortoise != hare {
		tortoise = tortoise.Next
		hare = hare.Next
	}

	return hare
}
