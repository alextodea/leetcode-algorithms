package strings

import (
	"testing"
)

// should return "12.345"
func TestThousandsSeparator(t *testing.T) {
	testInput := TargetInteger{value: 12345}
	assertionResult := "12.345"
	result, err := testInput.thousandsSeparator()

	if err != nil {
		t.Fail()
	}

	for i := range assertionResult {
		if assertionResult[i] != result[i] {
			t.Fail()
		}
	}
}

func TestThousandsSeparator2(t *testing.T) {
	testInput := TargetInteger{value:123456789}
	assertionResult := "123.456.789"
	result, err := testInput.thousandsSeparator()

	if err != nil {
		t.Fail()
	}

	for i := range assertionResult {
		if assertionResult[i] != result[i] {
			t.Fail()
		}
	}
}
