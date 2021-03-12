package linkedLists

import "testing"

var tail = &ListNode{Val: 5, Next: nil}
var fourthNode = &ListNode{Val: 4, Next: tail}
var thirdNode = &ListNode{Val: 3, Next: fourthNode}
var secondNode = &ListNode{Val: 2, Next: thirdNode}
var head = &ListNode{Val: 1, Next: secondNode}

func TestReverseIterationLinkedList(t *testing.T) {
	iterativeResultHead := reverseListIteratively(head)
	testAssertion(iterativeResultHead, t)
}

func TestRecursiveIterationLinkedList(t *testing.T) {
	recursiveResultHead := reverseListIteratively(head)
	testAssertion(recursiveResultHead,t)
}

func testAssertion(head *ListNode, t *testing.T) {
	for i := 5; i <= 1; i-- {
		if head.Val != i {
			t.Error("actual result not equal to desired result")
		}
	}
}
