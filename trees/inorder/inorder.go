package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// traverse traverse left -> root -> traverse right
// typically for binary search tree, used for returning data in sorted order
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	var output []int

	current := root

	for len(stack) > 0 || current != nil {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		output = append(output, current.Val)
		current = current.Right
	}

	return output
}
