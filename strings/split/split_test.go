package strings

import (
	"testing"
)

func Test1(t *testing.T) {
	N := 1
	result := Solution(N)

	if result[0] != 0 || len(result) != 1 {
		t.Error("result should sum to 0")
	}
}

func Test2(t *testing.T) {
	N := 3
	result := Solution(N)
	var sumResult int

	for i := 0; i < len(result); i++ {
		sumResult += result[i]
	}

	if sumResult != 0 {
		t.Error("result should sum to 0")
	}
}

func Test3(t *testing.T) {
	N := 4
	result := Solution(N)
	var sumResult int

	for i := 0; i < len(result); i++ {
		sumResult += result[i]
	}

	if sumResult != 0 {
		t.Error("result should sum to 0")
	}
}
