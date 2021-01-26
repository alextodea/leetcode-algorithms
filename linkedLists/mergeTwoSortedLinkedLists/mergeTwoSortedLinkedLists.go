package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) headValue() int {
	return l.Val
}

func (l *ListNode) tail() *ListNode {
	return l.Next
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.headValue() <= l2.headValue() {
		return &ListNode{Val: l1.headValue(), Next: merge(l1.tail(), l2)}
	}

	return &ListNode{Val: l2.headValue(), Next: merge(l1, l2.tail())}
}
