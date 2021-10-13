package strings

func Solution(N int) []int {
	output := []int{}

	for i := 0; i < N; i++ {
		output = append(output, i * 2 - N + 1)
	}

	return output
}