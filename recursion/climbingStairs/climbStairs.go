package recursion

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 2

	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
