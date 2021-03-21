package linkedLists

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// func flatten(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}

// 	var stack []*Node
// 	stack = append(stack, root)
// 	ll := &Node{Val: -1, Prev: nil, Next: nil, Child: nil}
// 	dummyNode := ll

// 	for len(stack) > 0 {
// 		stackNodeElement := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		ll.Next = &Node{Val: stackNodeElement.Val, Prev: ll, Next: nil, Child: nil}
// 		ll = ll.Next

// 		if stackNodeElement.Next != nil {
// 			stack = append(stack, stackNodeElement.Next)
// 			stackNodeElement.Next = nil
// 		}

// 		if stackNodeElement.Child != nil {
// 			stack = append(stack, stackNodeElement.Child)
// 			stackNodeElement.Child = nil
// 		}
// 	}

// 	output := dummyNode.Next
// 	output.Prev = nil
// 	return output
// }

func flatten(root *Node) *Node {
	if root == nil || root.Next == nil {
		return root
	}

	return flattenDepthFirstSearch(root, root.Next)
}

func flattenDepthFirstSearch(prev, curr *Node) *Node {
	if prev == nil || curr == nil {
		return nil
	}

	if curr.Child != nil {
		fromCurrentToChildrenNodes := flattenDepthFirstSearch(curr, curr.Child)
		return flattenDepthFirstSearch(fromCurrentToChildrenNodes, curr.Next)
	}

	return &Node{Val: prev.Val, Prev: prev.Prev, Next: flattenDepthFirstSearch(curr, curr.Next), Child: nil}
}
