package recursion

func kthGrammar(N int, K int) int {
	if N == 0 {
		return 0
	}

	pivot := K
	if pivot%2 != 0 {
		pivot++
	}

	x := kthGrammar(N-1, pivot/2)

	if K % 2 != 0 {
		return x
	}

	return (x+1) % 2
}
