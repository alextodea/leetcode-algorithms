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
