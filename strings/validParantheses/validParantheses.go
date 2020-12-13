package strings

type ValidParanthesesInput struct {
	value string
}

func (s *ValidParanthesesInput) areParanthesesValid() bool {
	if len(s.value)%2 != 0 {
		return false
	}

	var paranthesesMapping = map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	var stack = []string{}

	for i := 0; i < len(s.value); i++ {
		if string(s.value[i]) == "{" || string(s.value[i]) == "[" || string(s.value[i]) == "(" {
			stack = append(stack, paranthesesMapping[string(s.value[i])])
		} else if len(stack) > 0 && stack[len(stack)-1] == string(s.value[i]) {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
