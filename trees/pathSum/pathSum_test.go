package trees

import "testing"

func Constructor(val int) *TreeNode {
	if val == -1 {
		return nil
	}
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func TestPathSumWhenInputIsNull(t *testing.T) {
	if hasPathSum(nil, 1) != false {
		t.Error("result should be false")
	}
}

func TestPathSumWhenListHasOneNodeAndShouldReturnFalse(t *testing.T) {
	head := Constructor(2)

	if hasPathSum(head, 1) != false {
		t.Error("result should be false")
	}
}

func TestPathSumWhenListHasOneNodeAndShouldReturnTrue(t *testing.T) {
	head := Constructor(1)

	if hasPathSum(head, 1) != true {
		t.Error("result should be true")
	}
}

func TestPathSumWhenListHasMultipleNodesAndShouldReturnTrue(t *testing.T) {
	head := Constructor(1)
	head.Left = Constructor(2)
	head.Right = Constructor(2)
	head.Left.Left = Constructor(3)
	head.Left.Right = Constructor(4)
	head.Right.Left = Constructor(4)
	head.Right.Right = Constructor(5)

	if hasPathSum(head, 8) != true {
		t.Error("result should be true")
	}
}

func TestPathSumWhenListHasTwoNodesAndShouldReturnFalse(t *testing.T) {
	head := Constructor(1)
	head.Right = Constructor(2)

	if hasPathSum(head, 2) != false {
		t.Error("result should be false")
	}
}
