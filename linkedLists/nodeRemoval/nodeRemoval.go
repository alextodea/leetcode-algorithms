package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1st approach
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	dummyNode := &ListNode{Val: 0, Next: head}
	nthHead := fastForwardHeadNthTimes(dummyNode, n)
	remove(dummyNode, nthHead)
	return dummyNode.Next
}

func remove(dummyNode, nthHead *ListNode) {
	if nthHead.Next == nil {
		dummyNode.Next = dummyNode.Next.Next
		return
	}

	remove(dummyNode.Next, nthHead.Next)
}

func fastForwardHeadNthTimes(dummyNode *ListNode, n int) *ListNode {
	if n == 0 {
		return dummyNode
	}

	n--
	return fastForwardHeadNthTimes(dummyNode.Next, n)
}

// 2nd approach
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return nil
// 	}

// 	dummyNode := &ListNode{Val: 0, Next: head}
// 	listLength := getListLength(head)
// 	stopCounter := listLength - n
// 	var counter int
// 	remove(dummyNode, counter, stopCounter)

// 	return dummyNode.Next
// }

// func getListLength(head *ListNode) int {
// 	var length int

// 	for head != nil {
// 		length++
// 		head = head.Next
// 	}

// 	return length
// }

// func remove(dummyNode *ListNode, counter, stopCounter int) {
// 	if counter == stopCounter {
// 		dummyNode.Next = dummyNode.Next.Next
// 		return
// 	}
// 	counter++
// 	remove(dummyNode.Next, counter, stopCounter)
// }
