package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. iterative
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	sentinelNode := &ListNode{Val: -1, Next: head}
// 	return iterativeSwap(sentinelNode)
// }

// func iterativeSwap(head *ListNode) *ListNode {
// 	prevNode := head
// 	var firstNode *ListNode
// 	var secondNode *ListNode

// 	for head != nil && head.Next != nil {
// 		firstNode = head
// 		secondNode = head.Next

// 		prevNode.Next = secondNode
// 		firstNode.Next = secondNode.Next
// 		secondNode.Next = firstNode

// 		prevNode = firstNode
// 		head = firstNode.Next
// 	}

// 	return head.Next
// }

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	swappedTail := swapPairs(head.Next.Next)
	newFirstNode := head.Next
	newFirstNode.Next = head
	newFirstNode.Next.Next = swappedTail

	return newFirstNode
}
