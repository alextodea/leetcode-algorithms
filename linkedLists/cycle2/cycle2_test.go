package linkedLists

import "testing"

func TestCycleWhenListIsNull(t *testing.T) {
	result := detectCycle(nil)

	if result != nil {
		t.Error("result should be nil")
	}
}

func TestCycleWhenListHasOnlyOneNode(t *testing.T) {
	node1 := &ListNode{Val: 1, Next: nil}
	result := detectCycle(node1)

	if result != nil {
		t.Error("result should be nil")
	}
}

func TestCycleWhenListHasTwoNodesNoCycle(t *testing.T) {
	node2 := &ListNode{Val: 2, Next: nil}
	node1 := &ListNode{Val: 1, Next: node2}

	result := detectCycle(node1)

	if result != nil {
		t.Error("result should be nil")
	}
}

func TestCycleWhenListHasTwoNodesWithCycle(t *testing.T) {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: node1}
	node2.Next = node1

	result := detectCycle(node1)

	if result != nil {
		t.Error("result should be nil")
	}
}

func TestCycleWhenListHasFourNodesAndTailConnectsToSecondNode(t *testing.T) {
	node4 := &ListNode{Val: -4, Next: nil}
	node3 := &ListNode{Val: 0, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 3, Next: node2}
	node4.Next = node2

	result := detectCycle(node1)

	if result != node2 {
		t.Error("result should be node 2")
	}
}
