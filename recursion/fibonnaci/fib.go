package recursion

var cache = make(map[int]int)

func fib(n int) int {
	if n < 2 {
		return n
	}

	if _, nExistsInCache := cache[n]; !nExistsInCache {
		cache[n] = fib(n-1) + fib(n-2)
	}

	return cache[n]
}
