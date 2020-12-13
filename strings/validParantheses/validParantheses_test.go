package strings

import "testing"

func TestValidParantheses1(t *testing.T) {
	input := ValidParanthesesInput{value: "{[]}"}
	output := input.areParanthesesValid()

	if output == false {
		t.Error()
	}
}

func TestValidParantheses2(t *testing.T) {
	input := ValidParanthesesInput{value: "([)]"}
	output := input.areParanthesesValid()

	if output == true {
		t.Error()
	}
}

func TestValidParantheses3(t *testing.T) {
	input := ValidParanthesesInput{value: "(]"}
	output := input.areParanthesesValid()

	if output == true {
		t.Error()
	}
}

func TestValidParantheses4(t *testing.T) {
	input := ValidParanthesesInput{value: "()[]{}"}
	output := input.areParanthesesValid()

	if output == false {
		t.Error()
	}
}
