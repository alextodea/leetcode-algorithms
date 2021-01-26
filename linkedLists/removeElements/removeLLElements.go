package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	sentinel := &ListNode{Val: -1, Next: head}
	return removeTargetNode(sentinel, val)
}

func removeTargetNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	curr := removeTargetNode(head.Next, val)

	if curr == nil {
		return head
	}

	if (curr.Val == val && head.Val == -1) || (curr.Val == val && head.Val == -1 && curr.Next == nil) {
		curr = curr.Next
		return curr
	}

	if curr.Val == val && curr.Next == nil {
		head.Next = nil
		return head
	}

	if curr.Val == val {
		head.Next = curr.Next
		return head
	}

	if head.Val == -1 {
		return curr
	}

	return head
}
