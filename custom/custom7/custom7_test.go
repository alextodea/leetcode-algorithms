package custom

import "testing"

func TestWhenInputIsEmptyString(t *testing.T) {
	result := solution(123456)

	if result != 699 {
		t.Error("res should be 15")
	}
}
