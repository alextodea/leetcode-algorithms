package trees

import "testing"

func Constructor(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}



func TestWhenInputArrayHas1Element(t *testing.T) {
	result := buildTree([]int{-1}, []int{-1})
	expectedResult := Constructor(-1)
	if result.Val != expectedResult.Val {
		t.Error("result should be -1")
	}
}

func TestWhen4Nodes(t *testing.T) {
	root := Constructor(3)
	root.Left = Constructor(9)
	root.Right = Constructor(20)
	root.Right.Left = Constructor(15)
	root.Right.Right = Constructor(7)

	result := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})

	if result != nil {
		t.Error("sasa")
	}
}

func TestWhen5Nodes(t *testing.T) {
	root := Constructor(3)
	root.Left = Constructor(2)
	root.Left.Left = Constructor(5)
	root.Right = Constructor(1)
	root.Right.Left = Constructor(7)

	result := buildTree([]int{5, 2, 3, 7, 1}, []int{5, 2, 7, 1, 3})

	if result != nil {
		t.Error("sasa")
	}
}

func TestWhen2Nodes(t *testing.T) {
	root := Constructor(3)
	root.Left = Constructor(2)
	root.Left.Left = Constructor(5)
	root.Right = Constructor(1)
	root.Right.Left = Constructor(7)

	result := buildTree([]int{2, 1}, []int{2, 1})

	if result != nil {
		t.Error("sasa")
	}
}

func TestWhen4ConsecutiveValuesNodes(t *testing.T) {
	root := Constructor(3)
	root.Left = Constructor(2)
	root.Left.Left = Constructor(1)
	root.Right = Constructor(4)

	result := buildTree([]int{1, 2, 3, 4}, []int{1, 2, 4, 3})

	if result != nil {
		t.Error("sasa")
	}
}
