package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var nodeStack []*TreeNode
	var sumsStack []int
	var remainingSum int

	nodeStack = append(nodeStack, root)
	sumsStack = append(sumsStack, targetSum-root.Val)

	for len(nodeStack) > 0 {
		root = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		remainingSum = sumsStack[len(sumsStack)-1]
		sumsStack = sumsStack[:len(sumsStack)-1]

		if root.Left == nil && root.Right == nil && remainingSum == 0 {
			return true
		}

		if root.Right != nil {
			nodeStack = append(nodeStack, root.Right)
			sumsStack = append(sumsStack, remainingSum-root.Right.Val)
		}

		if root.Left != nil {
			nodeStack = append(nodeStack, root.Left)
			sumsStack = append(sumsStack, remainingSum-root.Left.Val)
		}
	}
	return false
}

// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	targetSum -= root.Val

// 	if root.Left == nil && root.Right == nil && targetSum == 0 {
// 		return true
// 	}

// 	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
// }
