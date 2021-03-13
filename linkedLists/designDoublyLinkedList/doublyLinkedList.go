package linkedLists

type MyLinkedList struct {
	head *Node
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: nil}
}

// delete node at given index
func (l *MyLinkedList) DeleteAtIndex(index int) {
	targetNode := l.GetNode(index)

	if targetNode == nil {
		return
	}

	prev := targetNode.Prev
	next := targetNode.Next

	if prev != nil {
		prev.Next = next
	} else {
		l.head = next
	}

	if next != nil {
		next.Prev = prev
	}
}

// add node a given index. if index == list length then append at tail, else if bigger return
func (l *MyLinkedList) AddAtIndex(index, value int) {
	if index == 0 {
		l.AddAtHead(value)
		return
	}

	prev := l.GetNode(index - 1)

	if prev == nil {
		return
	}

	newNode := &Node{Val: value, Next: prev.Next, Prev: prev}
	prev.Next = newNode

	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}
}

// add node at head
func (l *MyLinkedList) AddAtHead(value int) {
	l.head = &Node{Val: value, Next: l.head, Prev: nil}

	if l.head.Next != nil {
		l.head.Next.Prev = l.head
	}

	return
}

// add node at tail
func (l *MyLinkedList) AddAtTail(value int) {
	tail := l.GetTail()

	if tail == nil {
		l.AddAtHead(value)
		return
	}

	tail.Next = &Node{Val: value, Next: nil, Prev: tail}
	return
}

// returns tail node
func (l *MyLinkedList) GetTail() *Node {
	head := l.head
	for head != nil && head.Next != nil {
		head = head.Next
	}
	return head
}

// returns node at given index. if index is bigger than list length, null will be returned.
func (l *MyLinkedList) GetNode(index int) *Node {
	head := l.head
	for i := 0; i < index && head != nil; i++ {
		head = head.Next
	}
	return head
}

// returns node value at given index.
func (l *MyLinkedList) Get(index int) int {
	nodeAtTargetIndex := l.GetNode(index)

	if nodeAtTargetIndex != nil {
		return nodeAtTargetIndex.Val
	}

	return -1
}
