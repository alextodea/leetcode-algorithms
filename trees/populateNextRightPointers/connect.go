package trees

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	output := root

	traverseRecursive(root)

	return output
}

func traverseRecursive(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	connect(root.Left)
	connect(root.Next)

	return root
}

// func connect(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}

// 	leftmost := root
// 	var head *Node

// 	for leftmost.Left != nil {
// 		head = leftmost
// 		for head != nil {
//             head.Left.Next = head.Right
//             if head.Next != nil {
//                 head.Right.Next = head.Next.Left
//             }
//             head = head.Next
// 		}
// 		leftmost = leftmost.Left
// 	}
// 	return root
// }
