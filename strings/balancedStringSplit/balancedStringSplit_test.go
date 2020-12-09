package strings

import "testing"

func TestBalancedStringSplit1(t *testing.T) {
	testInput := TargetString{value: "RLRRRLLRLL"}
	testOutput, err := testInput.balancedStringSplit()

	if err != nil {
		t.Fail()
	}

	assertionOutput := 2

	if testOutput != assertionOutput {
		t.Fail()
	}
}

func TestBalancedStringSplit2(t *testing.T) {
	testInput := TargetString{value: "RLLLLRRRLR"}
	testOutput, err := testInput.balancedStringSplit()

	if err != nil {
		t.Fail()
	}

	assertionOutput := 3

	if testOutput != assertionOutput {
		t.Fail()
	}
}

func TestBalancedStringSplit3(t *testing.T) {
	testInput := TargetString{value: "LLLLRRRR"}
	testOutput, err := testInput.balancedStringSplit()

	if err != nil {
		t.Fail()
	}

	assertionOutput := 1

	if testOutput != assertionOutput {
		t.Fail()
	}
}

func TestBalancedStringSplit4(t *testing.T) {
	testInput := TargetString{value: "RLRRLLRLRL"}
	testOutput, err := testInput.balancedStringSplit()

	if err != nil {
		t.Fail()
	}

	assertionOutput := 4

	if testOutput != assertionOutput {
		t.Fail()
	}
}
