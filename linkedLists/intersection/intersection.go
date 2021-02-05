package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// method 1: using two pointers

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	if headA == nil || headB == nil {
// 		return nil
// 	}

// 	return compareNodes(headA, headB)
// }

// func compareNodes(headA, headB *ListNode) *ListNode {
// 	p1 := headA
// 	p2 := headB

// 	for p1 != p2 {
// 		p1 = p1.Next
// 		p2 = p2.Next

// 		if p1 == p2 {
// 			return p1
// 		}

// 		if p1 == nil {
// 			p1 = headB
// 		}

// 		if p2 == nil {
// 			p2 = headA
// 		}
// 	}

// 	return p1
// }

// method 2: using hashmpap
var mappings = map[*ListNode]ListNode{}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	headAMappings := storeNodesInMap(headA)

	for headB != nil {
		if _, nodesAreIntersected := headAMappings[headB]; nodesAreIntersected {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

func storeNodesInMap(head *ListNode) map[*ListNode]ListNode {

	if head == nil {
		return mappings
	}

	_, mappingExists := mappings[head]

	if !mappingExists {
		mappings[head] = *head
	}

	return storeNodesInMap(head.Next)
}
