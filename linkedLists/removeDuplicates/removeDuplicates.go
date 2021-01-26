package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		head = removeDuplicates(head.Next)
	}

	return &ListNode{Val: head.Val, Next: removeDuplicates(head.Next)}
}

// hashmap approach

// var nodesMap = map[int]int{}

// func removeDuplicates(head *ListNode) *ListNode {

// 	if head == nil {
// 		return nil
// 	}

// 	if _,ok := nodesMap[head.Val]; ok {
// 		return removeDuplicates(head.Next)
// 	}

// 	nodesMap[head.Val]++

// 	return &ListNode{Val: head.Val, Next: removeDuplicates(head.Next)}
// }
