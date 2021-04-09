package recursion

import "testing"

func compareLists(t *testing.T, l1, l2 []int) {
	if len(l1) == 0 || len(l2) == 0 || len(l1) != len(l2) {
		t.Error("list(s) is empty, or inequal")
	}

	for i:=0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			t.Error("list values are different")
		}
	}
}

func Test1(t *testing.T) {
	result := getRow(0)
	expectedResult := []int{1}

	compareLists(t,result,expectedResult)
}

func Test2(t *testing.T) {
	result := getRow(3)
	expectedResult := []int{1,3,3,1}

	compareLists(t,result,expectedResult)
}

func Test3(t *testing.T) {
	result := getRow(2)
	expectedResult := []int{1,2,1}

	compareLists(t,result,expectedResult)
}

func Test4(t *testing.T) {
	result := getRow(30)
	expectedResult := []int{1,2,1}

	compareLists(t,result,expectedResult)
}