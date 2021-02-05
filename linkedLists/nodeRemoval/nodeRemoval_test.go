package linkedLists

import "testing"

func areListsEqual(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}

	if l1 == nil || l2 == nil {
		return false
	}

	if l1.Val == l2.Val {
		return areListsEqual(l1.Next, l2.Next)
	}

	return false
}

func convertArrayToList(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}

	return &ListNode{Val: arr[0], Next: convertArrayToList(arr[1:])}
}

func TestNodeRemovalWhenListHasNoNodes(t *testing.T) {
	var head *ListNode
	result := removeNthFromEnd(head, 1)
	if result != nil {
		t.Error("result should be null")
	}
}

func TestNodeRemovalWhenListHasOneNode(t *testing.T) {
	head := convertArrayToList([]int{1})
	result := removeNthFromEnd(head, 1)
	if result != nil {
		t.Error("result should be null")
	}
}

func TestNodeRemovalWhenListHasTwoNodes(t *testing.T) {
	head := convertArrayToList([]int{1, 2})
	expectedResult := convertArrayToList([]int{2})
	result := removeNthFromEnd(head, 2)
	if areListsEqual(expectedResult, result) != true {
		t.Error("lists should be equal")
	}
}

func TestNodeRemovalWhenListHasMultipleNodes(t *testing.T) {
	head := convertArrayToList([]int{1, 2, 3, 4, 5})
	expectedResult := convertArrayToList([]int{1, 2, 3, 5})
	result := removeNthFromEnd(head, 2)
	if areListsEqual(expectedResult, result) != true {
		t.Error("lists should be equal")
	}
}
