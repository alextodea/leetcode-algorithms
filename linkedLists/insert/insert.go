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

func insertRecursive(slow, fast *Node, x int) {
	if slow.Val <= x && x <= slow.Next.Val {
		slow.Next = &Node{Val: x, Next: slow.Next}
		return
	}

	if fast.Next.Next == slow.Next { // next iteration is end of cycle
		if x <= slow.Next.Val {
			slow.Next.Next = &Node{Val: x, Next: slow.Next.Next}
		} else {
			slow.Next = &Node{Val: x, Next: slow.Next}
		}
		return
	}

	insertRecursive(slow.Next, fast.Next.Next, x)
}
