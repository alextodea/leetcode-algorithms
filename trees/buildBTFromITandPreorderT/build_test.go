package trees

import "testing"

func Test1(t *testing.T) {
	root := Constructor(3)
	root.Left = Constructor(9)
	root.Right = Constructor(20)
	root.Right.Left = Constructor(15)
	root.Right.Right = Constructor(7)

	result := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	if result != nil {
		t.Error("sdds")
	}
}
