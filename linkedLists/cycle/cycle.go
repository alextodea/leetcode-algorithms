package linkedLists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var nodesMap = map[*ListNode]int{}

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	if _, ok := nodesMap[head]; ok {
		fmt.Println(nodesMap[head])
		return true
	}

	nodesMap[head] = head.Val

	return hasCycle(head.Next)
}

func hasCycleTwoPointers(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	return determineCycle(head, head.Next)
}

func determineCycle(slow, fast *ListNode) bool {
	if fast.Next == nil || fast.Next.Next == nil {
		return false
	}

	if slow == fast {
		return true
	}

	return determineCycle(slow.Next, fast.Next.Next)
}
