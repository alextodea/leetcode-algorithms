package linkedLists

import "testing"

func convertArrayToList(arr []int) *Node {
	if len(arr) < 1 {
		return nil
	}

	head := &Node{Val: arr[0], Prev: nil, Next: convertArrayToList(arr[1:]), Child: nil}
	tail := head.Next

	if tail != nil {
		tail.Prev = head
	}
	return head
}

func testNodeLists(t *testing.T, expectedResult *Node, testResult *Node) {
	for testResult != nil {
		if testResult.Val != expectedResult.Val {
			t.Error("heads not equal")
		}
		testResult = testResult.Next
		expectedResult = expectedResult.Next
	}
}

func TestWhenListIsNull(t *testing.T) {
	var root *Node
	result := flatten(root)
	if result != nil {
		t.Error("result should be null")
	}
}

func TestWhenListNotNullCase1(t *testing.T) {
	root := convertArrayToList([]int{1, 2, 3, 4, 5, 6})
	rootChild := convertArrayToList([]int{7, 8, 9, 10})
	secondChild := convertArrayToList([]int{11, 12})
	rootChild.Next.Child = secondChild
	root.Next.Next.Child = rootChild
	result := flatten(root)
	expectedResult := convertArrayToList([]int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6})
	testNodeLists(t, result, expectedResult)
}
