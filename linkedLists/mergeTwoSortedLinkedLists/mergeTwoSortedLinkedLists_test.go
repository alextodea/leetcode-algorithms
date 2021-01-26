package linkedLists

import "testing"

var l1 = convertArrayToList([]int{1, 2, 3})
var l2 = convertArrayToList([]int{1, 2, 3, 4})
var expectedResult = convertArrayToList([]int{1, 1, 2, 2, 3, 3, 4})

func convertArrayToList(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}

	return &ListNode{Val: arr[0], Next: convertArrayToList(arr[1:])}
}

func TestMergeTwoSortedLinkedListsBothNil(t *testing.T) {
	testResult := merge(nil, nil)

	if testResult != nil {
		t.Error("result not ")
	}
}

func TestMergeTwoSortedLinkedListsL1Nil(t *testing.T) {
	testResult := merge(l1, nil)

	if testResult != l1 {
		t.Error("result not l1")
	}
}

func TestMergeTwoSortedLinkedListsL2Nil(t *testing.T) {
	testResult := merge(nil, l2)

	if testResult != l2 {
		t.Error("result not l2")
	}
}
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

func TestMergeTwoSortedLinkedLists(t *testing.T) {
	testResult := merge(l1, l2)

	if !areListsEqual(testResult, expectedResult) {
		t.Error("lists not equal")
	}
}
