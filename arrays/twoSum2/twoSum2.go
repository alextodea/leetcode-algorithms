package arrays

func twoSum2(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}

	var left int
	right := len(numbers) - 1
	var sum int

	for left < right {
		sum = numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if target > sum {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}
