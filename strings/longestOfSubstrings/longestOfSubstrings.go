package strings

var longestSSLength int

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return longestSSLength
	}

	return findLengthOfLS(s, map[string]bool{})
}

func findLengthOfLS(s string, seen map[string]bool) int {
	if len(s) == 0 {
		return longestSSLength
	}

	var thisSSLength int

	for i, r := range s {
		char := string(r)
		if seen[char] {
			if longestSSLength < thisSSLength {
				longestSSLength = thisSSLength
			}
			return findLengthOfLS(s[i:], map[string]bool{})
		}
		seen[char] = true
		thisSSLength++
	}

	return longestSSLength
}

// func lengthOfLongestSubstring(s string) int {

// 	var max int
// 	for i := 0; i < len(s); i++ {
// 		thisUniqueSubstringLength := getMaxSubstringLength(s[i:])
// 		if thisUniqueSubstringLength > max {
// 			max = thisUniqueSubstringLength
// 		}
// 	}
// 	return max
// }

// func getMaxSubstringLength(s string) int {
// 	var max int
// 	uniqueSymbols := make(map[string]bool)

// 	for i := 0; i < len(s); i++ {
// 		symbol := string(s[i])
// 		if _, exists := uniqueSymbols[symbol]; exists {
// 			return max
// 		}
// 		uniqueSymbols[symbol] = true
// 		max++
// 	}
// 	return max
// }
