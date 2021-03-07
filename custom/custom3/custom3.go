package custom

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int, B []int, N int) int {
	num_connected_cities := len(A)
	connected_cities := make([]int, N+1)
	var max_network_rank int

	for i := 0; i < num_connected_cities; i++ {
		connected_cities[A[i]]++
	}

	for i := 0; i < num_connected_cities; i++ {
		max_network_rank = max(max_network_rank, connected_cities[A[i]]+connected_cities[B[i]-1])
		fmt.Print(i)
	}

	return max_network_rank
}
