package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nrUnivalTrees int

func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countUnivalSubtrees(root.Left)
	right := countUnivalSubtrees(root.Right)

	if left == 0 && right == 0 {
		nrUnivalTrees++
		return nrUnivalTrees
	}

	if (left != 0 && root.Val != root.Left.Val) || (right != 0 && root.Val != root.Right.Val) {
		return nrUnivalTrees
	}

	nrUnivalTrees++

	return nrUnivalTrees
}
