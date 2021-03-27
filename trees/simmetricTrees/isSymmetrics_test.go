package trees

import "testing"

func Constructor(val int) *TreeNode {
	if val == -1 {
		return nil
	}
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func TestIsSymmetricsWhenInputNull(t *testing.T) {
	if isSymmetric(nil) != false {
		t.Error("res should be null")
	}
}

func TestIsSymmetricsWhenOnlyRoot(t *testing.T) {
	if isSymmetric(Constructor(1)) != true {
		t.Error("res should be true")
	}
}

func TestIsSymmetricsWhenIsSymmetric(t *testing.T) {

	root := Constructor(1)
	root.Left = Constructor(2)
	root.Right = Constructor(2)
	root.Left.Left = Constructor(3)
	root.Left.Right = Constructor(4)
	root.Right.Left = Constructor(4)
	root.Right.Right = Constructor(3)

	if isSymmetric(root) != true {
		t.Error("res should be true")
	}
}

func TestIsSymmetricsWhenIsNOTSymmetric(t *testing.T) {

	root := Constructor(1)
	root.Left = Constructor(2)
	root.Right = Constructor(2)
	root.Left.Right = Constructor(3)
	root.Right.Right = Constructor(3)

	if isSymmetric(root) != false {
		t.Error("res should be false")
	}
}

func TestIsSymmetricsWhenIsNOTSymmetric2(t *testing.T) {

	root := Constructor(1)
	root.Left = Constructor(2)
	root.Right = Constructor(2)
	root.Left.Left = Constructor(-1)
	root.Left.Right = Constructor(3)
	root.Right.Left = Constructor(3)
	root.Right.Right = Constructor(-1)
	root.Left.Right.Left = Constructor(-1)
	root.Left.Right.Right = Constructor(4)
	root.Right.Left.Left = Constructor(-1)
	root.Right.Left.Right = Constructor(4)

	if isSymmetric(root) != false {
		t.Error("res should be false")
	}
}
