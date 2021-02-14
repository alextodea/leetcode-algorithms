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

func TestRotateRightWhenListIsNull(t *testing.T) {
	var head *ListNode
	result := rotateRight(head, 1)
	if result != nil {
		t.Error("result should be null")
	}
}

func TestRotateRightWhenKIsSmallerThanNrNodes(t *testing.T) {
	head := convertArrayToList([]int{1, 2, 3, 4, 5})
	expectedResult := convertArrayToList([]int{4, 5, 1, 2, 3})
	result := rotateRight(head, 2)
	testNodeLists(t, expectedResult, result)
}

func TestRotateRightWhenKIsBiggerThanNrNodes(t *testing.T) {
	head := convertArrayToList([]int{0, 1, 2})
	expectedResult := convertArrayToList([]int{2, 0, 1})
	result := rotateRight(head, 4)
	testNodeLists(t, expectedResult, result)
}
