package linkedLists

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}


// O(n) with O(n)
var nodesmap = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	newHead := &Node{Val: head.Val}
	oldHead := head
	nodesmap[head] = newHead

	for head != nil {
		newHead.Next = addOrReturnCopiedNode(head.Next)
		newHead.Random = addOrReturnCopiedNode(head.Random)
		head = head.Next
		newHead = newHead.Next
	}

	return nodesmap[oldHead]
}

func addOrReturnCopiedNode(n *Node) *Node {
	if n == nil {
		return nil
	}

	if _, nodeExists := nodesmap[n]; !nodeExists {
		nodesmap[n] = &Node{Val: n.Val}
	}

	return nodesmap[n]
}
