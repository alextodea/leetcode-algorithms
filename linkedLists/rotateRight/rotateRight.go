package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	totalNrNodes := closeCircle(head)
	newHeadIndex := (totalNrNodes - k % totalNrNodes)
	return getNewHead(head, newHeadIndex)
}

func getNewHead(head *ListNode, newHeadIndex int) *ListNode {
	for i := 0; i < newHeadIndex-1; i++ {
		head = head.Next
	}

	newHead := head.Next
	head.Next = nil
	return newHead
}

func closeCircle(head *ListNode) int {
	node := head
	totalNrNodes := 1

	for node.Next != nil {
		node = node.Next
		totalNrNodes++
	}

	node.Next = head
	return totalNrNodes
}
