package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	dummyNode := &ListNode{Val: 101, Next: head}

	return deleteDuplicatesRecursively(dummyNode)
}

func deleteDuplicatesRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	middleNode := head.Next
	rightNode := head.Next.Next

	if rightNode == nil {
		if areNodesEqual(head, middleNode) {
			return deleteDuplicatesRecursively(middleNode)
		}
		return &ListNode{Val: middleNode.Val, Next: deleteDuplicatesRecursively(middleNode)}
	}

	if areNodesEqual(head, middleNode) || areNodesEqual(middleNode, rightNode) {
		return deleteDuplicatesRecursively(middleNode)
	}

	return &ListNode{Val: middleNode.Val, Next: deleteDuplicatesRecursively(middleNode)}
}

func areNodesEqual(n1, n2 *ListNode) bool {
	if n1.Val == n2.Val {
		return true
	}
	return false
}

// method 1
// var nodesMap = map[int]*ListNode{}

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	_, nodeExistsInMap := nodesMap[head.Val]

// 	if nodeExistsInMap {
// 		return deleteDuplicates(head.Next)
// 	}

// 	if head.Val == head.Next.Val {
// 		nodesMap[head.Val] = head
// 		return deleteDuplicates(head.Next.Next)
// 	}

// 	nodesMap[head.Val] = head

// 	node := &ListNode{Val: head.Val, Next: deleteDuplicates(head.Next)}
// 	return node
// }
