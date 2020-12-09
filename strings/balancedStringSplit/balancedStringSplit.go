package strings

import "errors"

type TargetString struct {
	value string
}

func (s *TargetString) balancedStringSplit() (int, error) {
	inputLength := len(s.value)

	if inputLength < 1 {
		return 0, errors.New("input string is empty")
	}

	nrBalancedStrings := 0

	for i := 0; i < inputLength; i++ {
		char := string(s.value[i])
		symbols[char]++

		if symbols["R"] == symbols["L"] {
			nrBalancedStrings++
			symbols["R"] = 0
			symbols["L"] = 0
		}
	}

	return nrBalancedStrings, nil

}

var symbols = map[string]int{
	"R": 0,
	"L": 0,
}
