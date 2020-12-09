package strings

import "testing"

func TestIsLongPressedName1(t *testing.T) {
	targetNames := TargetNames{name: "alex", typed: "aaleex"}
	testResult := targetNames.isLongPressed()

	if testResult == false {
		t.Error()
	}
}

func TestIsLongPressedName2(t *testing.T) {
	targetNames := TargetNames{name: "saeed", typed: "ssaaedd"}
	testResult := targetNames.isLongPressed()

	if testResult == true {
		t.Error()
	}
}

func TestIsLongPressedName3(t *testing.T) {
	targetNames := TargetNames{name: "leelee", typed: "lleeelee"}
	testResult := targetNames.isLongPressed()

	if testResult == false {
		t.Error()
	}
}

func TestIsLongPressedName4(t *testing.T) {
	targetNames := TargetNames{name: "laiden", typed: "laiden"}
	testResult := targetNames.isLongPressed()

	if testResult == false {
		t.Error()
	}
}

func TestIsLongPressedName5(t *testing.T) {
	targetNames := TargetNames{name: "alex", typed: "aaleexa"}
	testResult := targetNames.isLongPressed()

	if testResult == false {
		t.Error()
	}
}
