package arrays

func search(nums []int, target int) int {
	if nums == nil || len(nums) < 1 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	var pivot int

	for left <= right {
		pivot = left + ((right - left) / 2)
		if target == nums[pivot] {
			return pivot
		}
		if target > nums[pivot] {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return -1
}
