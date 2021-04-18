package recursion

import "testing"

func Test1(t *testing.T) {
	result := myPow(2.00, 0)

	if result != 1 {
		t.Error("res should be 1")
	}
}

func Test2(t *testing.T) {
	result := myPow(2.00, 1)

	if result != 2.00 {
		t.Error("res should be 2.00")
	}
}

func Test3(t *testing.T) {
	result := myPow(2.00000, 10)

	if result != 1024.00000 {
		t.Error("res should be 1024.00000")
	}
}

func Test4(t *testing.T) {
	result := myPow(2.00000, -2)

	if result != 0.25 {
		t.Error("res should be 0.25")
	}
}
