package linkedLists

import "testing"

func TestIsListPalindrome(t *testing.T) {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 2, Next: nil}
	node4 := &ListNode{Val: 1, Next: nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	if isPalindrome(node1) != true {
		t.Error("linked list is palindrome")
	}
}

func TestIsListPalindrome2(t *testing.T) {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}

	node1.Next = node2

	if isPalindrome(node1) != false {
		t.Error("linked list is palindrome")
	}
}

// test also with null and one single node
