package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var traversalOutput []int
	var stack []*TreeNode

	stack = append(stack, root)

	for len(stack) > 0 {
		stackLastElementIndex := len(stack) - 1
		node := stack[stackLastElementIndex]
		stack = stack[:stackLastElementIndex]
		traversalOutput = append(traversalOutput, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return traversalOutput
}

// var output []int

// func preorderTraversalRecursive(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}

// 	output = append(output, root.Val)
// 	preorderTraversalRecursive(root.Left)
// 	preorderTraversalRecursive(root.Right)

// 	return output
// }
