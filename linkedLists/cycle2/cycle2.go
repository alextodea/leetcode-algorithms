package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	nodeIntersection := determineNodeIntersection(head.Next, head.Next.Next)

	if nodeIntersection == nil {
		return nil
	}

	return determineCycleEntry(head, nodeIntersection)
}

func determineNodeIntersection(slow, fast *ListNode) *ListNode {
	if fast == nil || fast.Next == nil {
		return nil
	}

	if slow == fast {
		return slow
	}

	return determineNodeIntersection(slow.Next, fast.Next.Next)
}

func determineCycleEntry(head, nodeIntersection *ListNode) *ListNode {

	if head == nodeIntersection {
		return head
	}

	return determineCycleEntry(head.Next, nodeIntersection.Next)
}
