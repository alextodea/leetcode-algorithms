package arrays

// import "testing"

// func compareLists(t *testing.T, arr1, arr2 []int) {
// 	if len(arr1) != len(arr2) {
// 		t.Error("arrays are not equal")
// 	}
// 	for i := 0; i < len(arr1); i++ {
// 		if arr1[i] != arr2[i] {
// 			t.Error("array values are not identic")
// 		}
// 	}
// }

// func TestMoveZeroesWhenListIsNull(t *testing.T) {
// 	result := moveZeroes(nil)

// 	if result != nil {
// 		t.Error("result should be nil")
// 	}
// }

// func TestMoveZeroesWhenListHasOnlyOneElement(t *testing.T) {
// 	result := moveZeroes([]int{0})

// 	compareLists(t, result, []int{0})
// }

// func TestMoveZeroesWhenListHas5Elements(t *testing.T) {
// 	testArr := []int{0, 1, 0, 3, 12}
// 	result := moveZeroes(testArr)
// 	expectedOutput := []int{1, 3, 12, 0, 0}

// 	compareLists(t, result, expectedOutput)
// }

// func TestMoveZeroesWhenListHas3Elements(t *testing.T) {
// 	testArr := []int{0, 0, 1}
// 	result := moveZeroes(testArr)
// 	expectedOutput := []int{1, 0, 0}

// 	compareLists(t, result, expectedOutput)
// }
