package trees

import "testing"

func compareLists(t *testing.T, l1 []int, l2 []int) {
	if len(l1) != len(l2) {
		t.Error("lists are not equal")
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			t.Error("list elements are different")
		}
	}
}

func Constructor(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}
func TestMaxDepthWhenListIsNull(t *testing.T) {
	result := maxDepth(nil)

	if result != 0 {
		t.Error("res should be nil")
	}
}

func TestMaxDepthWhenListHasOnlyRoot(t *testing.T) {
	result := maxDepth(Constructor(1))

	if result != 1 {
		t.Error("res should be 1")
	}
}

func TestMaxDepthWhenListHasDepth3(t *testing.T) {
	root := Constructor(3)
	root.Left = Constructor(9)
	root.Right = Constructor(20)
	root.Right.Left = Constructor(15)
	root.Right.Right = Constructor(7)
	
	result := maxDepth(root)

	if result != 3 {
		t.Error("res should be 3")
	}
}
