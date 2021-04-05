package trees

import "testing"

func Constructor(val int) *Node {
	return &Node{Val: val, Left: nil, Right: nil, Next: nil}
}

func Test1(t *testing.T) {
	root := Constructor(1)
	root.Left = Constructor(2)
	root.Right = Constructor(3)
	root.Left.Left = Constructor(4)
	root.Left.Right = Constructor(5)
	root.Right.Left = Constructor(6)
	root.Right.Right = Constructor(7)

	result := connect(root)

	if result != nil {
		t.Error("sasa")
	}
}
