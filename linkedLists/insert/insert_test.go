package linkedLists

import "testing"

func convertArrayToList(arr []int) *Node {
	if len(arr) < 1 {
		return nil
	}

	list := &Node{Val: arr[0], Next: nil}
	head := list

	for i := 1; i < len(arr); i++ {
		list.Next = &Node{Val: arr[i], Next: nil}
		list = list.Next
	}

	list.Next = head

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
	x := 3
	result := insert(root, x)
	expectedResult := &Node{Val: x, Next: nil}
	expectedResult.Next = expectedResult
	if result != expectedResult {
		t.Error("result should be as expected")
	}
}

func TestWhenInputContainsListWithOneNode(t *testing.T) {
	list := convertArrayToList([]int{1})
	x := 0
	result := insert(list, x)
	expectedResult := convertArrayToList([]int{1, 0})

	testNodeLists(t, result, expectedResult)
}

func TestWhenInputContainsListWithMultipleNodes(t *testing.T) {
	list := convertArrayToList([]int{3, 4, 1})
	x := 2
	result := insert(list, x)
	expectedResult := convertArrayToList([]int{3, 4, 1, 2})

	testNodeLists(t, result, expectedResult)
}

func Test4(t *testing.T) {
	list := convertArrayToList([]int{3, 5, 1})
	x := 0
	result := insert(list, x)
	expectedResult := convertArrayToList([]int{3, 5, 0, 1})

	testNodeLists(t, result, expectedResult)
}

func Test5(t *testing.T) {
	list := convertArrayToList([]int{1, 3, 5})
	x := 0
	result := insert(list, x)
	expectedResult := convertArrayToList([]int{1, 3, 5, 0})

	testNodeLists(t, result, expectedResult)
}

func Test6(t *testing.T) {
	list := convertArrayToList([]int{3, 4, 1})
	x := 2
	result := insert(list, x)
	expectedResult := convertArrayToList([]int{3, 4, 1, 2})

	testNodeLists(t, result, expectedResult)
}

func Test7(t *testing.T) {
	list := convertArrayToList([]int{1, 3, 5})
	x := 1
	result := insert(list, x)
	expectedResult := convertArrayToList([]int{1, 1, 3, 5})

	testNodeLists(t, result, expectedResult)
}

func Test8(t *testing.T) {
	list := convertArrayToList([]int{1, 3, 5})
	x := 4
	result := insert(list, x)
	expectedResult := convertArrayToList([]int{1, 3, 4, 5})

	testNodeLists(t, result, expectedResult)
}

func Test9(t *testing.T) {
	list := convertArrayToList([]int{1, 3, 5})
	x := 6
	result := insert(list, x)
	expectedResult := convertArrayToList([]int{1, 3, 5, 6})

	testNodeLists(t, result, expectedResult)
}