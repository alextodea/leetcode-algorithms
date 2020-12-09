package strings

import (
	"fmt"
)

type TargetInteger struct {
	value int
}

// ThousandsSeparator separates thousands within an integer and returns a string
func (n *TargetInteger) thousandsSeparator() (string, error) {

	inputAsString := fmt.Sprint(n.value)
	inputLength := len(inputAsString)

	if inputLength < 1 {
		return "", fmt.Errorf("input string is empty")
	}

	var output string

	for i := inputLength - 1; i >= 0; i-- {
		output = string(inputAsString[i]) + output
		if (inputLength - i) % 3 == 0 && i > 0 {
			output = "." + output
		}
	}

	return output, nil
}
