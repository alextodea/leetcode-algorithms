package linkedLists

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	head        *Node
	numberNodes int
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: nil, numberNodes: 0}
}

// Get the value of the indexth node in the linked list. If the index is invalid, return -1.
func (ll *MyLinkedList) Get(index int) int {
	head := ll.head
	counter := 0

	for head != nil {
		if index == counter {
			return head.Val
		}

		counter++
		head = head.Next
	}

	return -1
}

func (ll *MyLinkedList) AddAtHead(val int) {
	newHead := &Node{Val: val, Next: ll.head}
	ll.head = newHead
}

// Append a node of value val as the last element of the linked list.
func (ll *MyLinkedList) AddAtTail(val int) {
	head := ll.head

	if head == nil {
		ll.AddAtHead(val)
		return
	}

	for head.Next != nil {
		head = head.Next
	}

	head.Next = &Node{Val: val, Next: nil}
}

// Add a node of value val before the indexth node in the linked list.
// If index equals the length of the linked list, the node will be appended to the end of the linked list.
// If index is greater than the length, the node will not be inserted.
func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		ll.AddAtHead(val)
		return
	}

	head := ll.head
	counter := 1

	for head != nil && head.Next != nil {
		if index == counter {
			temp := head.Next
			head.Next = &Node{Val: val, Next: temp}
			return
		}
		head = head.Next
		counter++
	}

	if index == counter {
		head.Next = &Node{Val: val, Next: nil}
		return
	}
	return
}

// Delete the indexth node in the linked list, if the index is valid.
func (ll *MyLinkedList) DeleteAtIndex(index int) {
	head := ll.head

	if ll.head == nil {
		return
	}

	if index == 0 {
		ll.head = ll.head.Next
		return
	}

	counter := 1

	for head != nil && head.Next != nil {
		if index == counter {
			head.Next = head.Next.Next
			return
		}
		head = head.Next
		counter++
	}
	return
}
