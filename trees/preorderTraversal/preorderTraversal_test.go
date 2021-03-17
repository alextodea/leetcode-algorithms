package trees

import "testing"

func compareLists(t *testing.T, l1, l2 []int) {
	if len(l1) != len(l2) {
		t.Error("lists not equal")
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			t.Error("values not equal")
		}
	}
}

func TestPreorderTraversal1(t *testing.T) {
	right1 := &TreeNode{Val: 2, Left: nil, Right: nil}
	right2 := &TreeNode{Val: 3, Left: nil, Right: nil}
	root := &TreeNode{Val: 1, Left: nil, Right: nil}

	root.Right = right1
	right1.Right = right2

	result := preorderTraversal(root)
	expectedOutput := []int{1, 2, 3}

	compareLists(t, result, expectedOutput)
}

func TestPreorderTraversal2(t *testing.T) {
	root := &TreeNode{Val: 1, Left: nil, Right: nil}

	result := preorderTraversal(root)
	expectedOutput := []int{1}

	compareLists(t, result, expectedOutput)
}
