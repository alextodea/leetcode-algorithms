package custom

import "strconv"

func solution(n int) int {
	if n == 0 {
		return 0
	}

	var digits []int
	for _, char := range strconv.Itoa(n) {
		digits = append(digits, int(char-'0'))
	}

	return productOfElements(digits) - sumOfElements(digits)
}

func sumOfElements(digits []int) int {
	var output int
	for _, digit := range digits {
		output += digit
	}

	return output
}

func productOfElements(digits []int) int {
	output := 1
	for _, digit := range digits {
		output = output * digit
	}

	return output
}
