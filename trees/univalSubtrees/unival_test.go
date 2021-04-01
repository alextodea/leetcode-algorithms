package trees

import "testing"

func Constructor(val int) *TreeNode {
	if val == -1 {
		return nil
	}
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func TestUnivalWhenInputIsNull(t *testing.T) {
	if countUnivalSubtrees(nil) != 0 {
		t.Error("result should be 0")
	}
}

func TestUnivalWhenTreeHas4UnivalSubtrees(t *testing.T) {
	root := Constructor(5)
	root.Left = Constructor(1)
	root.Left.Left = Constructor(5)
	root.Left.Right = Constructor(5)
	root.Right = Constructor(5)
	root.Right.Right = Constructor(5)

	result := countUnivalSubtrees(root)

	if result != 4 {
		t.Error("result should be 0")
	}
}

func TestUnivalWhenTreeHasNoInnerSubtreeBesidesLeafs(t *testing.T) {
	root := Constructor(1)
	root.Left = Constructor(2)
	root.Left.Left = Constructor(3)
	root.Left.Right = Constructor(4)
	root.Right = Constructor(5)
	root.Right.Right = Constructor(6)

	result := countUnivalSubtrees(root)

	if result != 3 {
		t.Error("result should be 2")
	}
}

func TestUnivalWhenTreeHas6UnivalSubtrees(t *testing.T) {
	root := Constructor(5)
	root.Left = Constructor(5)
	root.Left.Left = Constructor(5)
	root.Left.Right = Constructor(5)
	root.Right = Constructor(5)
	root.Right.Right = Constructor(5)

	result := countUnivalSubtrees(root)

	if result != 6 {
		t.Error("result should be 6")
	}
}
