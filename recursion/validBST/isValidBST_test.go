package recursion

import "testing"

func Test1(t *testing.T) {
	result := isValidBST(nil)

	if result != true {
		t.Error("res should be true")
	}
}

func Test2(t *testing.T) {
	root := Constructor(2)
	root.Left = Constructor(1)
	root.Right = Constructor(3)
	result := isValidBST(root)

	if result != true {
		t.Error("res should be true")
	}
}

func Test3(t *testing.T) {
	root := Constructor(5)
	rootLeft := Constructor(1)
	rootRight := Constructor(4)
	rootRightLeft := Constructor(3)
	rootRightRight := Constructor(6)

	root.Left = rootLeft
	root.Right = rootRight
	root.Right.Left = rootRightLeft
	root.Right.Right = rootRightRight


	result := isValidBST(root)

	if result != false {
		t.Error("res should be false")
	}
}