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

func TestPreorderTraversalWhenInputIsNull(t *testing.T) {
	result := preorderTraversal(nil)

	if result != nil {
		t.Error("res should be nil")
	}
}

func Constructor(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func TestPreorderTraversalWhenTreeHasTheRootOnly(t *testing.T) {
	root := Constructor(1)
	result := preorderTraversal(root)
	expectedResult := []int{1}

	compareLists(t, result, expectedResult)
}

func TestPreorderTraversalWhenTreeHas3Nodes(t *testing.T) {
	root := Constructor(1)
	right := Constructor(2)
	rightLeft := Constructor(3)

	root.Right = right
	right.Left = rightLeft
	result := preorderTraversal(root)
	expectedResult := []int{1, 2, 3}

	compareLists(t, result, expectedResult)
}

func TestPreorderTraversalWhenTreeHas2NodesBothDerivedFromRoot(t *testing.T) {
	root := Constructor(3)
	left := Constructor(1)
	right := Constructor(2)

	root.Right = right
	root.Left = left
	result := preorderTraversal(root)
	expectedResult := []int{3, 1, 2}

	compareLists(t, result, expectedResult)
}
