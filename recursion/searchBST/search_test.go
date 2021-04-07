package recursion

import "testing"

func Constructor(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func Test1(t *testing.T) {
	if searchBST(nil, 1) != nil {
		t.Error("res should be nil")
	}
}

func Test2(t *testing.T) {
	root := Constructor(4)
	root.Left = Constructor(2)
	root.Left.Left = Constructor(1)
	root.Left.Right = Constructor(3)
	root.Right = Constructor(7)

	result := searchBST(root, 5)

	if result != nil {
		t.Error("sasa")
	}
}
