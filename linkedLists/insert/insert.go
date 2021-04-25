package linkedLists

type Node struct {
	Val  int
	Next *Node
}

func Constructor(val int) *Node {
	return &Node{Val: val, Next: nil}
}

func insert(aNode *Node, x int) *Node {
	newNode := Constructor(x)
	if aNode == nil {
		newNode.Next = newNode
		return newNode
	}

	prev := aNode.Next
	curr := prev.Next
	var doInsert bool

	for prev != aNode {
		if prev.Val <= x && x <= curr.Val {
			doInsert = true
		}

		if prev.Val > curr.Val {
			if prev.Val <= x || x <= curr.Val {
				doInsert = true
			}
		}

		if doInsert == true {
			prev.Next = newNode
			newNode.Next = curr
			return aNode
		}

		prev = prev.Next
		curr = curr.Next
	}

	prev.Next = newNode
	newNode.Next = curr

	return aNode
}
