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

func TestPartitionWhenEmptyList(t *testing.T) {
	var head *ListNode
	result := partition(head, 1)

	if result != nil {
		t.Error("result should be null")
	}
}

func TestPartitionWhenListContainsOnlyTwoElements(t *testing.T) {
	head := convertArrayToList([]int{2, 1})
	expectedResult := convertArrayToList([]int{1, 2})
	result := partition(head, 2)

	if areListsEqual(result, expectedResult) != true {
		t.Error("lists should be equal")
	}
}

func TestPartitionWhenListContainsSeveralElements(t *testing.T) {
	head := convertArrayToList([]int{1, 4, 3, 2, 5, 2})
	expectedResult := convertArrayToList([]int{1, 2, 2, 4, 3, 5})
	result := partition(head, 3)

	if areListsEqual(result, expectedResult) != true {
		t.Error("lists should be equal")
	}
}
