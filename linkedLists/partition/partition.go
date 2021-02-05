package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func initializeList(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	headListBeforeX := initializeList(0)
	tailListBeforeX := headListBeforeX
	headListAfterX := initializeList(0)
	tailListAfterX := headListAfterX

	for head != nil {
		if head.Val < x {
			tailListBeforeX.Next = head
			tailListBeforeX = tailListBeforeX.Next
		} else {
			tailListAfterX.Next = head
			tailListAfterX = tailListAfterX.Next
		}
		head = head.Next
	}
	tailListAfterX.Next = nil
	tailListBeforeX.Next = headListAfterX.Next

	return headListBeforeX.Next
}
