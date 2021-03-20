package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var output []int

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		output = append(output, root.Val)
		return output
	}

	return traverse(root)
}

func traverse(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	output = append(output, root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)

	return output
}

// func preorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return []int{root.Val}
// 	}

// 	var output []int
// 	var stack []*TreeNode

// 	stack = append(stack, root)

// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		output = append(output, node.Val)

// 		if node.Right != nil {
// 			stack = append(stack, node.Right)
// 		}

// 		if node.Left != nil {
// 			stack = append(stack, node.Left)
// 		}
// 	}
// 	return output
// }
