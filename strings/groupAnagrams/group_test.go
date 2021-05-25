package strings

import "testing"

func Test1(t *testing.T) {
	result := groupAnagrams([]string{""})
	expectedResult := [][]string{{""}}
	if result[0][0] != expectedResult[0][0] {
		t.Error("result not correct")
	}
}

func Test2(t *testing.T) {
	result := groupAnagrams([]string{"a"})
	expectedResult := [][]string{{"a"}}
	if result[0][0] != expectedResult[0][0] {
		t.Error("result not correct")
	}
}

func Test3(t *testing.T) {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	expectedResult := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}

	for i, arr := range result {
		for j, anagram := range arr {
			if expectedResult[i][j] != anagram {
				t.Error("incorrect")
			}
		}
	}
}
