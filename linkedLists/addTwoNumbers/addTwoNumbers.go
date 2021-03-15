package linkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	currentNodes := doElementaryAddition(l1, l2, 0)

	return currentNodes
}

func doElementaryAddition(l1, l2 *ListNode, carryDigit int) *ListNode {

	if l1 == nil && l2 == nil && carryDigit == 1 {
		return doElementaryAddition(&ListNode{Val: 0, Next: nil}, &ListNode{Val: 0, Next: nil}, carryDigit)
	}

	dummyNode := &ListNode{Val: 0, Next: nil}

	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return doElementaryAddition(dummyNode, l2, carryDigit)
	}

	if l2 == nil {
		return doElementaryAddition(l1, dummyNode, carryDigit)
	}

	sum := l1.Val + l2.Val + carryDigit

	if sum > 9 {
		carryDigit = 1
		sum -= 10
	} else {
		carryDigit = 0
	}

	return &ListNode{Val: sum, Next: doElementaryAddition(l1.Next, l2.Next, carryDigit)}
}
