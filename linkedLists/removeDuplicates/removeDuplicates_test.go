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

func TestRemoveDuplicatesWhenListIsNull(t *testing.T) {
	var head *ListNode
	testResult := removeDuplicates(head)
	testNodeLists(t, head, testResult)
}

func TestRemoveDuplicatesWhenDuplicatesAtBeginningOfList(t *testing.T) {
	head := convertArrayToList([]int{1, 1, 2})
	expectedResult := convertArrayToList([]int{1, 2})
	testResult := removeDuplicates(head)
	testNodeLists(t, expectedResult, testResult)
}

func TestRemoveDuplicatesWhenMultipleDuplicates(t *testing.T) {
	head := convertArrayToList([]int{1, 1, 2, 3, 3})
	expectedResult := convertArrayToList([]int{1, 2, 3})
	testResult := removeDuplicates(head)
	testNodeLists(t, expectedResult, testResult)
}
