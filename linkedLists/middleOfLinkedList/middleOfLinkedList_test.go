package linkedLists

import "testing"

func TestMiddleOfLinkedList(t *testing.T) {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	middleNode := getMiddleNode(node1)

	if middleNode.Val != 3 {
		t.Error("Middle node should be 3")
	}
}

func TestMiddleOfLinkedList2(t *testing.T) {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	node6 := &ListNode{Val: 6, Next: nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	middleNode := getMiddleNode(node1)

	if middleNode.Val != 4 {
		t.Error("Middle node should be 4")
	}
}

func TestMiddleOfLinkedList3(t *testing.T) {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 2, Next: nil}
	node4 := &ListNode{Val: 1, Next: nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	middleNode := getMiddleNode(node1)

	if middleNode.Val != 2 {
		t.Error("Middle node should be 4")
	}
}

func TestMiddleOfLinkedList4(t *testing.T) {
	node1 := &ListNode{Val: 1, Next: nil}

	middleNode := getMiddleNode(node1)

	if middleNode.Val != 1 {
		t.Error("Middle node should be 4")
	}
}

func TestMiddleOfLinkedList5(t *testing.T) {

	if getMiddleNode(nil) != nil {
		t.Error("Middle node should be 4")
	}
}

// func TestMiddleOfLinkedList(t *testing.T) {
// 	node1 := &ListNode{Val: 1, Next: nil}
// 	node2 := &ListNode{Val: 2, Next: nil}
// 	node3 := &ListNode{Val: 3, Next: nil}
// 	node4 := &ListNode{Val: 4, Next: nil}
// 	node5 := &ListNode{Val: 5, Next: nil}

// 	node1.Next = node2
// 	node2.Next = node3
// 	node3.Next = node4
// 	node4.Next = node5

// 	middleNode := node1.getMiddleNode()

// 	if middleNode.Val != 3 {
// 		t.Error("Middle node should be 3")
// 	}
// }

// func TestMiddleOfLinkedList2(t *testing.T) {
// 	node1 := &ListNode{Val: 1, Next: nil}
// 	node2 := &ListNode{Val: 2, Next: nil}
// 	node3 := &ListNode{Val: 3, Next: nil}
// 	node4 := &ListNode{Val: 4, Next: nil}
// 	node5 := &ListNode{Val: 5, Next: nil}
// 	node6 := &ListNode{Val: 6, Next: nil}

// 	node1.Next = node2
// 	node2.Next = node3
// 	node3.Next = node4
// 	node4.Next = node5
// 	node5.Next = node6

// 	middleNode := node1.getMiddleNode()

// 	if middleNode.Val != 4 {
// 		t.Error("Middle node should be 4")
// 	}
// }
