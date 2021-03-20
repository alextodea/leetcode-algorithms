package custom

import "sort"

func Solution(A []int) bool {
	if A == nil {
		return false
	}

	if len(A)%2 != 0 {
		return false
	}

	sort.Ints(A)

	for i := 1; i < len(A); i += 2 {
		if A[i-1] != A[i] {
			return false
		}
	}
	return true
}
