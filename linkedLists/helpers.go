package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConvertArrayToList(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}

	return &ListNode{Val: arr[0], Next: ConvertArrayToList(arr[1:])}
}
