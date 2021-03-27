package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return traverse(root, root)
}

func traverse(t1, t2 *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, t1)
	queue = append(queue, t2)

	for len(queue) > 0 {
		t1 = queue[0]
		t2 = queue[1]
		queue = queue[2:]

		if t1 == nil && t2 == nil {
			continue
		}

		if t1 == nil || t2 == nil || (t1.Val != t2.Val) {
			return false
		}

		queue = append(queue, t1.Left)
		queue = append(queue, t2.Right)

		queue = append(queue, t1.Right)
		queue = append(queue, t2.Left)
	}

	return true
}

// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return false
// 	}
// 	return isSymmetricRecursive(root, root)
// }

// func isSymmetricRecursive(left, right *TreeNode) bool {
// 	if left == nil && right == nil {
// 		return true
// 	}

// 	if left == nil || right == nil {
// 		return false
// 	}

// 	if left.Val != right.Val {
// 		return false
// 	}

// 	return isSymmetricRecursive(left.Left, right.Right) && isSymmetricRecursive(left.Right, right.Left)

// }
