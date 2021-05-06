package recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Constructor(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func isValidBST(root *TreeNode) bool {
	return traverseTreeRecursively(root, nil, nil)
}

func traverseTreeRecursively(root, lowerBound, upperBound *TreeNode) bool {
	if root == nil {
		return true
	}

	if upperBound != nil && root.Val >= upperBound.Val {
		return false
	}

	if lowerBound != nil && root.Val <= lowerBound.Val {
		return false
	}

	if (root.Left != nil && root.Left.Val > root.Val) || (root.Right != nil && root.Right.Val <= root.Val) {
		return false
	}

	return traverseTreeRecursively(root.Left, lowerBound, root) && traverseTreeRecursively(root.Right, root, upperBound)
}
