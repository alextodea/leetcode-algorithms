package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NodeConstructor(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

var inorderNodesIndexesMaps = make(map[int]int)
var thisPostorder []int
var postOrderLastElementIndex int

func buildTree(inorder, postorder []int) *TreeNode {
	for i, val := range inorder {
		inorderNodesIndexesMaps[val] = i
	}

	thisPostorder = postorder
	postOrderLastElementIndex = len(thisPostorder) - 1

	return buildTreeRecursively(0, postOrderLastElementIndex)
}

func buildTreeRecursively(left, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootValue := thisPostorder[postOrderLastElementIndex]
	postOrderLastElementIndex--

	index := inorderNodesIndexesMaps[rootValue]
	root := NodeConstructor(rootValue)
	root.Right = buildTreeRecursively(index+1, right)
	root.Left = buildTreeRecursively(left, index-1)
	return root
}
