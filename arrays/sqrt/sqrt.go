package arrays

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left := 2
	right := x / 2
	var pivot int
	var num int

	for left <= right {
		pivot = left + ((right - left)/2)
		num = pivot * pivot
		if x < num {
			right = pivot - 1
		} else if x > num {
			left = pivot + 1
		} else {
			return pivot
		}

	}
	return right
}
