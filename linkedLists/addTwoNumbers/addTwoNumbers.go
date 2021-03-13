package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	currentNodes := addTwoNumbers(l1.Next, l2.Next)

	return currentNodes
}
