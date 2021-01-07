package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMiddleNode(head *ListNode) *ListNode {
	return getMiddleNodeRecursive(head, head)
}

func getMiddleNodeRecursive(slow *ListNode, fast *ListNode) *ListNode {
	if fast == nil {
		return slow
	}

	if fast.Next == nil {
		return slow
	}

	return getMiddleNodeRecursive(slow.Next, fast.Next.Next)
}

func (l *ListNode) getMiddleNode() *ListNode {
	counter := 1
	node := l

	for node.Next != nil {
		node = node.Next
		counter++
	}

	i := 0

	counter = counter / 2

	for i != counter {
		l = l.Next
		i++
	}

	return l
}
