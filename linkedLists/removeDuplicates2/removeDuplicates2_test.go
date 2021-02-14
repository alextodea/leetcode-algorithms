package linkedLists

import "testing"

func testNodeLists(t *testing.T, expectedResult *ListNode, testResult *ListNode) {
	for testResult != nil {
		if testResult.Val != expectedResult.Val {
			t.Error("heads not equal")
		}
		testResult = testResult.Next
		expectedResult = expectedResult.Next
	}
}

func convertArrayToList(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}

	return &ListNode{Val: arr[0], Next: convertArrayToList(arr[1:])}
}

func TestRemoveDuplicates2WhenListIsNull(t *testing.T) {
	var head *ListNode
	result := deleteDuplicates(head)

	if result != nil {
		t.Error("result should be null")
	}
}

func TestRemoveDuplicates2WhenListHasDuplicates1(t *testing.T) {
	head := convertArrayToList([]int{1, 1, 1, 2, 3})
	expectedResult := convertArrayToList([]int{2, 3})
	result := deleteDuplicates(head)

	testNodeLists(t, expectedResult, result)
}

func TestRemoveDuplicates2WhenListHasDuplicates2(t *testing.T) {
	head := convertArrayToList([]int{1, 2, 3, 3, 4, 4, 5})
	expectedResult := convertArrayToList([]int{1, 2, 5})
	result := deleteDuplicates(head)

	testNodeLists(t, expectedResult, result)
}

func TestRemoveDuplicates2WhenListHasDuplicates3(t *testing.T) {
	head := convertArrayToList([]int{1, 1, 1, 2, 3})
	expectedResult := convertArrayToList([]int{2, 3})
	result := deleteDuplicates(head)

	testNodeLists(t, expectedResult, result)
}
