package linkedLists

import "testing"

func TestIntersectionWhenBothListsAreNull(t *testing.T) {
	var headA *ListNode
	var headB *ListNode
	result := getIntersectionNode(headA, headB)

	if result != nil {
		t.Error("result should be nil")
	}
}

func TestIntersectionWhenOneListIsNull(t *testing.T) {
	headA := &ListNode{Val: 1, Next: nil}
	var headB *ListNode
	result := getIntersectionNode(headA, headB)
	if result != nil {
		t.Error("result should be nil")
	}
}
func TestIntersectionWhenListsDoNotIntesect(t *testing.T) {
	headA := &ListNode{Val: 1, Next: nil}
	headA.Next = &ListNode{Val: 2, Next: nil}
	headB := &ListNode{Val: 3, Next: nil}
	headB.Next = &ListNode{Val: 4, Next: nil}
	result := getIntersectionNode(headA, headB)
	if result != nil {
		t.Error("result should be nil")
	}
}
func TestIntersectionWhenListsIntersect(t *testing.T) {
	headC := &ListNode{Val: 4, Next: nil}
	headA := &ListNode{Val: 1, Next: headC}
	headB := &ListNode{Val: 2, Next: headC}

	result := getIntersectionNode(headA, headB)
	if result != headC {
		t.Error("result should be headC")
	}
}
