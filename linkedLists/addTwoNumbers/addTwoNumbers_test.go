package linkedLists

import "testing"

func TestAddTwoNumbersWhenInputsAreNil(t *testing.T) {
	result := addTwoNumbers(nil, nil)
	if result != nil {
		t.Error("result should be null")
	}
}

func TestAddTwoNumbersWhenOneInputIsNil(t *testing.T) {
	result := addTwoNumbers(nil, convertArrayToList([]int{1}))
	if result != nil {
		t.Error("result should be null")
	}
}

func TestAddTwoNumbersWhenBothListsAreEqual(t *testing.T) {
	l1 := convertArrayToList([]int{2, 4, 3})
	l2 := convertArrayToList([]int{5, 6, 4})
	result := addTwoNumbers(l1, l2)
	expectedResult := convertArrayToList([]int{7, 0, 8})

	areListsEqual(result, expectedResult)
}

func TestAddTwoNumbersWhenListsAreUnequal(t *testing.T) {
	l1 := convertArrayToList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := convertArrayToList([]int{8, 9, 9, 9})
	result := addTwoNumbers(l1, l2)
	expectedResult := convertArrayToList([]int{8, 9, 9, 9, 0, 0, 0, 1})

	areListsEqual(result, expectedResult)
}

func TestAddTwoNumbersWhenListsAreUnequal2(t *testing.T) {
	l2 := convertArrayToList([]int{9, 9, 9, 9, 9, 9, 9})
	l1 := convertArrayToList([]int{9, 9, 9, 9})
	result := addTwoNumbers(l1, l2)
	expectedResult := convertArrayToList([]int{8, 9, 9, 9, 0, 0, 0, 1})

	areListsEqual(result, expectedResult)
}

///////////////////////// HELPERS
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
