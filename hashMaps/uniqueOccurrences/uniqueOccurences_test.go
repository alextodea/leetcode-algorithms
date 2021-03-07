package hashMaps

import "testing"

func TestUniqueOccurencesWhenInputArrHasLength1(t *testing.T) {
	testArr := []int{1}
	result := uniqueOccurrences(testArr)

	if result != true {
		t.Error("resould should be true")
	}
}

func TestUniqueOccurencesWhenInputArrHasUniqueOccurences(t *testing.T) {
	testArr := []int{1, 2, 2, 1, 1, 3}
	result := uniqueOccurrences(testArr)

	if result != true {
		t.Error("resould should be true")
	}
}

func TestUniqueOccurencesWhenInputArrHasTwoEqualOccurences(t *testing.T) {
	testArr := []int{1, 2}
	result := uniqueOccurrences(testArr)

	if result != false {
		t.Error("resould should be false")
	}
}

func TestUniqueOccurencesWhenInputArrHasUniqueOccurencesWithNegativeIntegers(t *testing.T) {
	testArr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	result := uniqueOccurrences(testArr)

	if result != true {
		t.Error("resould should be true")
	}
}
