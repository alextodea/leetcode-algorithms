package strings

import "fmt"

type TargetNames struct {
	name  string
	typed string
}

func (n *TargetNames) isLongPressed() bool {
	nameChars := make(map[string]int, len(n.name)-1)
	typedChars := make(map[string]int, len(n.typed)-1)

	for _, char := range n.name {
		_, ok := nameChars[string(char)]

		if !ok {
			nameChars[string(char)] = 1
		} else {
			nameChars[string(char)]++
		}
	}

	for _, char := range n.typed {
		_, ok := typedChars[string(char)]

		if !ok {
			typedChars[string(char)] = 1
		} else {
			typedChars[string(char)]++
		}

		fmt.Println(string(char), nameChars[string(char)])
		fmt.Println(string(char), typedChars[string(char)])
		fmt.Print("")
	}

	for _, char := range n.name {
		nrOccurencesInTyped, ok := typedChars[string(char)]

		if len(nameChars) != len(typedChars) {
			return false
		}

		// if char from typed does not occur in name
		if !ok || (nameChars[string(char)] > nrOccurencesInTyped) {
			return false
		}
	}

	return true
}
