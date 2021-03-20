package custom

import "testing"

func TestWhenInputIsEmptyString(t *testing.T) {
	result := Solution("")

	if result != 0 {
		t.Error("res should be 0")
	}
}

func TestWhenInputIsOneCharString(t *testing.T) {
	result := Solution("s")

	if result != 0 {
		t.Error("res should be 0")
	}
}

func TestWhenInputIsTwoCharsString(t *testing.T) {
	result := Solution("sa")

	if result != 0 {
		t.Error("res should be 0")
	}
}

func TestWhenInputIsThreeCharsStringThatContains2Digrams(t *testing.T) {
	result := Solution("aaa")

	if result != 1 {
		t.Error("res should be 1")
	}
}

func Test3(t *testing.T) {
	result := Solution("aakmaakmakda")

	if result != 7 {
		t.Error("res should be 1")
	}
}

// func TestWhenArrayContainsTwoDistinctValues(t *testing.T) {
// 	result := Solution([]int{1, 2})

// 	if result != false {
// 		t.Error("res should be nil")
// 	}
// }

// func TestWhenArrayContainsTwoPairs(t *testing.T) {
// 	result := Solution([]int{1, 2, 2, 1})

// 	if result != true {
// 		t.Error("res should be true")
// 	}
// }

// func TestWhenArrayContainsThreeOfAKind(t *testing.T) {
// 	result := Solution([]int{7, 7, 7})

// 	if result != false {
// 		t.Error("res should be nil")
// 	}
// }

// func TestWhenArrayContainsOnlyOnePair(t *testing.T) {
// 	result := Solution([]int{1, 2, 2, 3})

// 	if result != false {
// 		t.Error("res should be nil")
// 	}
// }
