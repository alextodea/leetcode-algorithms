package recursion

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	return powIterative(x, n)
}

func powIterative(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	ans := 1.00
	current_product := x

	for i := n; i > 0; i /= 2 {
		if i%2 != 0 {
			ans = ans * current_product
		}
		current_product = current_product * current_product
	}
	return ans
}
