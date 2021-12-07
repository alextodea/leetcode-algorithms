package recursion

type Result struct {
	combinations [][]int
	seen         map[int]bool
	n            int
	k            int
}

func (r *Result) backtrack(start int) [][]int {
	if start > r.n {
		return r.combinations
	}

	combination := []int{}

	for i := start; i <= r.n; i++ {
		if len(combination) == r.k {
			r.combinations = append(r.combinations, combination)
			combination = []int{}
		}
		combination = append(combination, i)
	}
	next := start + 1
	return r.backtrack(next)
}

func combine(n int, k int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	result := Result{combinations: [][]int{}, n: n, k: k}

	return result.backtrack(1)
}
