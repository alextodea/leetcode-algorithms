package strings

import "strconv"

type TargetStrings struct {
	s1 string
	s2 string
}

type AddResult struct {
	result string
	carry  int
}

func (s *TargetStrings) add() AddResult {

	if len(s.s1) < 1 {
		return AddResult{result: "", carry: 0}
	}

	if len(s.s1) != len(s.s2) {
		s.equalizeStrings()
	}

	tailStrings := TargetStrings{s1: s.s1[1:], s2: s.s2[1:]}
	addedTailStrings := tailStrings.add()
	digit1 := int(rune(s.s1[0] - '0'))
	digit2 := int(rune(s.s2[0] - '0'))
	sum := digit1 + digit2 + addedTailStrings.carry

	return AddResult{result: strconv.Itoa(sum%10) + addedTailStrings.result, carry: sum / 10}
}

func (s *TargetStrings) equalizeStrings() TargetStrings {

	if len(s.s1) > len(s.s2) {
		for i := 0; i < len(s.s1)-len(s.s2); i++ {
			s.s2 = "0" + s.s2
		}
	}

	if len(s.s1) < len(s.s2) {
		for i := 0; i < len(s.s2)-len(s.s1); i++ {
			s.s1 = "0" + s.s1
		}
	}

	return TargetStrings{s1: s.s1, s2: s.s2}
}