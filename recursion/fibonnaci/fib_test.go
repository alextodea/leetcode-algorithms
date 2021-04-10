package recursion

import "testing"

func Test1(t *testing.T) {
	if fib(0) != 0 {
		t.Error("res should be 0")
	}
}

func Test2(t *testing.T) {
	if fib(1) != 1 {
		t.Error("res should be 0")
	}
}

func Test3(t *testing.T) {
	if fib(3) != 2 {
		t.Error("res should be 2")
	}
}

func Test4(t *testing.T) {
	if fib(4) != 3 {
		t.Error("res should be 2")
	}
}
