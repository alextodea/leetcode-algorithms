package strings

import (
	"testing"
)

func Test(t *testing.T) {
	testInput := "abcabcbb"
	assertionResult := 3
	result := lengthOfLongestSubstring(testInput)

	if result != assertionResult {
		t.Error("res should be 3")
	}
}

func Test2(t *testing.T) {
	testInput := "bbbbb"
	assertionResult := 1
	result := lengthOfLongestSubstring(testInput)

	if result != assertionResult {
		t.Error("res should be 1")
	}
}

func Test3(t *testing.T) {
	testInput := "pwwkew"
	assertionResult := 3
	result := lengthOfLongestSubstring(testInput)

	if result != assertionResult {
		t.Error("res should be 3")
	}
}
