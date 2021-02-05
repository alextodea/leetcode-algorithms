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

func TestSwapNodesEmptyList(t *testing.T) {
	var head *ListNode
	result := swapPairs(head)

	if result != nil {
		t.Error("result should be null")
	}
}

func TestSwapNodes(t *testing.T) {
	head := convertArrayToList([]int{1,2,3,4})
	 expectedResult := convertArrayToList([]int{2,1,4,3})
	 result := swapPairs(head)

	 if !areListsEqual(result,expectedResult) {
		 t.Error("nodes are not swapped correctly")
	 }
}