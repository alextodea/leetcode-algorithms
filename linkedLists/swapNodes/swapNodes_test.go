package linkedLists

import "testing"

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

	 if result != expectedResult {
		 t.Error("nodes are not swapped correctly")
	 }
}