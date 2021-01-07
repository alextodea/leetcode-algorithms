package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func reverseListIteratively(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var oldFirst *ListNode
	var first = head

	for first != nil {
		temp := first.Next
		first.Next = oldFirst
		oldFirst = first
		first = temp
	}

	return oldFirst
}
