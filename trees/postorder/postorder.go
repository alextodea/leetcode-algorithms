package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	var output []int
	stack = append(stack, root)

	for len(stack) > 0 {
		stackLastIndex := len(stack) - 1
		current := stack[stackLastIndex]
		stack = stack[:stackLastIndex]
		output = append(output, current.Val)

		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}

	return reverseInts(output)
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
