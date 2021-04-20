package recursion

import "testing"

func Test1(t *testing.T) {
	result := kthGrammar(1,1)

	if result != 0 {
		t.Error("res should be 0")
	}
}

func Test2(t *testing.T) {
	result := kthGrammar(2,1)

	if result != 0 {
		t.Error("res should be 0")
	}
}

func Test3(t *testing.T) {
	result := kthGrammar(2,2)

	if result != 1 {
		t.Error("res should be 1")
	}
}