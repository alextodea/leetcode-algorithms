package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return recursiveMaxDepth(root, 1)
}

func recursiveMaxDepth(node *TreeNode, depth int) int {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return depth
	}

	depth++

	leftDepth := recursiveMaxDepth(node.Left, depth)
	rightDepth := recursiveMaxDepth(node.Right, depth)

	if leftDepth >= rightDepth {
		return leftDepth
	}

	return rightDepth
}
