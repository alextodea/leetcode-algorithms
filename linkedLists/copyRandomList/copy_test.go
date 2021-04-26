package linkedLists

import "testing"

func Test1(t *testing.T) {
	result := copyRandomList(nil)
	if result != nil {
		t.Error("res should be nil")
	}
}

func Test2(t *testing.T) {
	result := copyRandomList(&Node{Val: 1, Next: nil, Random: nil})
	if result.Val != 1 {
		t.Error("res should be 1")
	}
}

func Test3(t *testing.T) {
	n2 := &Node{Val: 13}
	n3 := &Node{Val: 11}
	n4 := &Node{Val: 10}
	n5 := &Node{Val: 1}

	head := &Node{Val: 7}
	head.Next = n2
	head.Random = nil

	n2.Next = n3
	n2.Random = head

	n3.Next = n4
	n3.Random = n5

	n4.Next = n5
	n4.Random = n3

	n5.Next = nil
	n5.Random = head
	
	result := copyRandomList(head)
	if result.Val != 1 {
		t.Error("res should be 1")
	}
}
