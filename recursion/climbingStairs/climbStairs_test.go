package recursion

import "testing"

func Test1(t *testing.T) {
	result := climbStairs(1)

	if result != 1 {
		t.Error("res should be 1")
	}
}

func Test2(t *testing.T) {
	result := climbStairs(2)

	if result != 2 {
		t.Error("res should be 2")
	}
}

func Test3(t *testing.T) {
	result := climbStairs(3)

	if result != 3 {
		t.Error("res should be 3")
	}
}

func Test4(t *testing.T) {
	result := climbStairs(4)

	if result != 5 {
		t.Error("res should be 5")
	}
}
