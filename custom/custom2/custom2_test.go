package custom

import "testing"

var testArr1 = []int{4, 8, 12, 16}
var testArr2 = []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
var testArr3 = []int{100}

func TestCustom1(t *testing.T) {
	result := Solution(testArr1)
	expected := 2

	if result != expected {
		t.Error("fial")
	}
}

func TestCustom2(t *testing.T) {
	result := Solution(testArr2)
	expected := 5

	if result != expected {
		t.Error("fial")
	}
}

func TestCustom3(t *testing.T) {
	result := Solution(testArr3)
	expected := 1

	if result != expected {
		t.Error("fial")
	}
}
