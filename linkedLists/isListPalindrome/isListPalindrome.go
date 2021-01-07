package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

var test ListNode

func isPalindrome(l *ListNode) bool {
	return areListsEqual(l, reverseList(l, nil))
}

func reverseList(l *ListNode, acc *ListNode) *ListNode {
	if l == nil {
		return acc
	}

	return reverseList(l.Next, &ListNode{Val: l.Val, Next: acc})
}

func areListsEqual(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}

	if l1 == nil || l2 == nil {
		return false
	}

	if l1.Val == l2.Val {
		return areListsEqual(l1.Next, l2.Next)
	}

	return false
}
