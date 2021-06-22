package recursion

import (
	"testing"
)

func Test1(t *testing.T) {
	result := multiply("10", "10")

	if result != "100" {
		println(result)
		t.Error("res should be 100")
	}
}

func Test2(t *testing.T) {
	result := multiply("3141592653589793238462643383279502884197169399375105820974944592", "2718281828459045235360287471352662497757247093699959574966967627")

	if result != "56088" {
		println(result)
		t.Error("res should be 56088")
	}
}

// func Test3(t *testing.T) {
// 	result := multiply("123", "456")

// 	if result != "56088" {
// 		println(result)
// 		t.Error("res should be 56088")
// 	}
// }

// func Test4(t *testing.T) {
// 	result := multiply("5678", "1234")

// 	if result != "56088" {
// 		println(result)
// 		t.Error("res should be 56088")
// 	}
// }
