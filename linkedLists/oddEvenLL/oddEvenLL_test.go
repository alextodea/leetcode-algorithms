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

func TestOddEvenLLWhenListIsNull(t *testing.T) {
	if oddEvenList(nil) != nil {
		t.Error("result should be nil")
	}
}

func TestOddEvenLLWhenListHasSingleElement(t *testing.T) {
	head := convertArrayToList([]int{1})
	result := oddEvenList(head)
	if result != head {
		t.Error("result should be nil")
	}
}

func TestOddEvenLLWhenListHas3Elements(t *testing.T) {
	head := convertArrayToList([]int{1, 2, 3})
	expectedResult := convertArrayToList([]int{1, 3, 2})
	result := oddEvenList(head)

	areListsEqual(expectedResult, result)
}

func TestOddEvenLLWhenListHas5Elements(t *testing.T) {
	head := convertArrayToList([]int{1, 2, 3, 4, 5})
	expectedResult := convertArrayToList([]int{1, 3, 5, 2, 4})
	result := oddEvenList(head)

	areListsEqual(expectedResult, result)
}
