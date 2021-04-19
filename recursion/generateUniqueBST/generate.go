package recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Constructor(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func generateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{Constructor(1)}
	}

	return generateBSTRecursively(1, n)
}

func generateBSTRecursively(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	list := []*TreeNode{}

	for i := start; i <= end; i++ {
		left := generateBSTRecursively(start, i-1)
		right := generateBSTRecursively(i+1, end)

		for _, l := range left {
			for _, r := range right {
				node := Constructor(i)
				node.Left = l
				node.Right = r
				list = append(list, node)
			}
		}
	}
	return list
}
