package linkedLists

import "testing"

var head = convertArrayToList([]int{1, 2, 3, 4})

func convertArrayToList(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}

	return &ListNode{Val: arr[0], Next: convertArrayToList(arr[1:])}
}

func TestRemoveLLElementsWhenFirstNodeIsTargetValue(t *testing.T) {
	expectedHead := convertArrayToList([]int{2, 3, 4})
	testVal := 1
	resultHead := removeElements(head, testVal)

	if resultHead != expectedHead {
		t.Error("remove first list node test has failed")
	}
}

func TestRemoveLLElementsWhenLasttNodeIsTargetValue(t *testing.T) {
	expectedHead := convertArrayToList([]int{1,2, 3})
	testVal := 4
	resultHead := removeElements(head, testVal)

	if resultHead != expectedHead {
		t.Error("remove first list node test has failed")
	}
}

func TestRemoveLLElementsWhenRandomIsTargetValue(t *testing.T) {
	expectedHead := convertArrayToList([]int{1,2,4})
	testVal := 3
	resultHead := removeElements(head, testVal)

	if resultHead != expectedHead {
		t.Error("remove first list node test has failed")
	}
}

func TestRemoveLLElementsWhenHeadIsNill(t *testing.T) {
	var expectedHead *ListNode
	expectedHead = nil
	testVal := 1
	resultHead := removeElements(expectedHead, testVal)

	if resultHead != expectedHead {
		t.Error("remove first list node test has failed")
	}
}

func TestRemoveLLElementsWhenHeadIsTargetAndListHasOnlyOneElement(t *testing.T) {
	expectedHead := convertArrayToList([]int{1})
	oneNodeListHead := &ListNode{Val: 1, Next: nil}
	testVal := 1
	resultHead := removeElements(oneNodeListHead, testVal)

	if resultHead != expectedHead {
		t.Error("remove first list node test has failed")
	}
}

// fails with input [1] 1