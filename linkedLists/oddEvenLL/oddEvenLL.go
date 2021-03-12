package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odds := groupNodes(head)
	evens := groupNodes(head.Next)

	oddsHeader := odds

	for odds.Next != nil {
		odds = odds.Next
	}

	odds.Next = evens

	return oddsHeader
}

func groupNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return &ListNode{Val: head.Val, Next: groupNodes(head.Next.Next)}
}
