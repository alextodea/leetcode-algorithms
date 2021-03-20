package strings

func lengthOfLongestSubstring(s string) int {

	var max int
	for i := 0; i < len(s); i++ {
		thisUniqueSubstringLength := getMaxSubstringLength(s[i:])
		if thisUniqueSubstringLength > max {
			max = thisUniqueSubstringLength
		}
	}
	return max
}

func getMaxSubstringLength(s string) int {
	var max int
	uniqueSymbols := make(map[string]bool)

	for i := 0; i < len(s); i++ {
		symbol := string(s[i])
		if _, exists := uniqueSymbols[symbol]; exists {
			return max
		}
		uniqueSymbols[symbol] = true
		max++
	}
	return max
}
