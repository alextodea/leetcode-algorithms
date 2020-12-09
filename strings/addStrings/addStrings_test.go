package strings

import "testing"

func TestAddStrings(t *testing.T) {
	targetStrings := TargetStrings{s1: "123", s2: "4567"}
	result := "4690"
	testResult := targetStrings.add()

	for i := range testResult.result {
		if result[i] != testResult.result[i] {
			t.Error()
		}
	}
}
