package linkedLists

import "testing"

func convertArrayToList(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}

	return &ListNode{Val: arr[0], Next: convertArrayToList(arr[1:])}
}

func TestCyclicLLWhenHeadIsNull(t *testing.T) {
	var head *ListNode
	testResult := hasCycle(head)
	if testResult == true {
		t.Error("null head should be false")
	}
}

func TestCyclicLLWhenCyclicWithTwoElements(t *testing.T) {
	tail := &ListNode{}
	head := &ListNode{}

	head.Val = 1
	head.Next = tail
	tail.Val = 2
	tail.Next = head

	testResult := hasCycle(head)

	if testResult == false {
		t.Error("null head should be true")
	}
}

func TestCyclicLLWhenCyclicFromLastToSecondElement(t *testing.T) {
	tail := &ListNode{}
	node3 := &ListNode{}
	node2 := &ListNode{}
	head := &ListNode{}

	head.Val = 3
	head.Next = node2
	node2.Val = 2
	node2.Next = node3
	node3.Val = 0
	node3.Next = tail
	tail.Val = 4
	tail.Next = node2

	testResult := hasCycle(head)

	if testResult == false {
		t.Error("null head should be true")
	}
}