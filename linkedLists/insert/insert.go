package linkedLists

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	newNode := &Node{Val: x, Next: nil}
	if aNode == nil {
		newNode.Next = newNode
		return newNode
	}

	if aNode.Next == nil {
		aNode.Next = newNode
		newNode.Next = aNode
		return aNode
	}

	head := aNode

	insertRecursive(aNode, aNode.Next, x)

	return head
}

func insertRecursive(curr, next *Node, x int) {
	if curr.Val <= x && x <= next.Val || *curr == *next {
		curr.Next = &Node{Val: x, Next: next}
		return
	}
	insertRecursive(next, next.Next, x)
}
