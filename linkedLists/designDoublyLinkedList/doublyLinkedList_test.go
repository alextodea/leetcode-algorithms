package linkedLists

import "testing"

func convertArrayToList(arr []int) MyLinkedList {
	l := Constructor()

	for _, val := range arr {
		l.AddAtHead(val)
	}

	return l
}

func testNodeLists(t *testing.T, expectedResult *MyLinkedList, testResult *MyLinkedList) {
	testResultHead := testResult.head
	expectedResultHead := expectedResult.head
	for testResultHead != nil {
		if testResultHead.Val != expectedResultHead.Val {
			t.Error("heads not equal")
		}
		testResultHead = testResultHead.Next
		expectedResultHead = expectedResultHead.Next
	}
}

func TestDesignLLAddAtHead(t *testing.T) {
	ll := Constructor()
	ll.AddAtHead(1)
	ll.AddAtHead(2)
	ll.AddAtHead(3)
	expectedResult := convertArrayToList([]int{1, 2, 3})
	testNodeLists(t, &ll, &expectedResult)
}

func TestDesignLLGetNode(t *testing.T) {
	ll := Constructor()
	ll.AddAtHead(1)
	ll.AddAtHead(2)
	firstIndexValue := ll.Get(0)

	if firstIndexValue != 2 {
		t.Error("first index value should be 2")
	}

	secondIndexValue := ll.Get(1)

	if secondIndexValue != 1 {
		t.Error("second index value should be 1")
	}

	indexValueBiggerThanListLength := ll.Get(3)

	if indexValueBiggerThanListLength != -1 {
		t.Error("returned value should be -1 because index is too big")
	}
}

func TestDesignLLAddAtTail(t *testing.T) {
	ll := Constructor() // nil

	// case 1: head is null before adding at tail, add at tail, return node val at index 0
	ll.AddAtTail(1) // 1 -> nil
	valAtFirstIndex := ll.Get(0)

	if valAtFirstIndex != 1 {
		t.Error("tail value should be 1")
	}

	// case 2: add an extra node at head, at another one at tail, then return tail
	ll.AddAtHead(2) // 2 -> 1 -> nil
	ll.AddAtTail(3) // 2 -> 1 -> 3 -> nil

	tailValue := ll.Get(2) // should return value of the third node in the list

	if tailValue != 3 {
		t.Error("tail value should be 3")
	}
}

func TestDesignLLAddAtIndex(t *testing.T) {
	ll := Constructor() // nil

	// case 1: ll has nil head, add at index 0 should add at head
	ll.AddAtIndex(0, 1) // 1 -> nil

	valueAtIndex0 := ll.Get(0) // should return "1"

	if valueAtIndex0 != 1 {
		t.Error("linked list's head value should be 1")
	}

	// case 2: index == list length
	ll.AddAtIndex(1, 2)        // 1 -> 2 | should append at tail because index == length of ll
	valueAtIndex1 := ll.Get(1) // should return 2

	if valueAtIndex1 != 2 {
		t.Error("value at tail should be 2")
	}

	// case 3: index > list length
	ll.AddAtIndex(3, 3) // should not add any new node because index > list length (which is 2)

	valueAtIndex2 := ll.Get(2) // should return -1 because ll.AddAtIndex(3,3) should not have added any new node

	if valueAtIndex2 != -1 {
		t.Error("should return -1")
	}
}

func TestDesignLLDeleteAtIndex(t *testing.T) {
	ll := Constructor() // nil

	// case 1: ll has nil head, deleteAtIndex 0 should not do anything
	ll.DeleteAtIndex(0)

	valueAtIndex0 := ll.Get(0) // should return -1

	if valueAtIndex0 != -1 {
		t.Error("linked list's GET method should return -1 because head is nil")
	}

	// case 2: add a node at head and delete that node
	ll.AddAtHead(1)
	ll.DeleteAtIndex(0)
	valueAtIndex0 = ll.Get(0) // should return -1

	if valueAtIndex0 != -1 {
		t.Error("linked list's GET method should return -1 because head is nil")
	}

	// case 3: add a node at head and attempt to delete node at an index that is bigger than list length
	ll.AddAtHead(2)
	ll.DeleteAtIndex(1)
	valueAtIndex0 = ll.Get(0)

	if valueAtIndex0 != 2 {
		t.Error("value at index 0 should be 2")
	}
}