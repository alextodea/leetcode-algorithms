package linkedLists

type ListNode struct {
	Val   int
	Prev  *ListNode
	Next  *ListNode
	Child *ListNode
}

func flatten(root *ListNode) *ListNode {
	if root == nil {
		return root
	}
	sentinelNode := &ListNode{Val: 0, Prev: nil, Next: root, Child: nil}
	root.Prev = sentinelNode
	flattenRecursively(sentinelNode, sentinelNode.Next)
	return sentinelNode.Next
}

func flattenRecursively(prev, curr *ListNode) *ListNode {
	if curr == nil {
		return prev
	}

	curr.Prev = prev
	prev.Next = curr

	tempNext := curr.Next
	tail := flattenRecursively(curr, curr.Child)
	curr.Child = nil

	return flattenRecursively(tail, tempNext)
}
