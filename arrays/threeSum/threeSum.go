package arrays

import "sort"

var output [][]int

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return output
	}

	sort.Ints(nums)

	for i := 1; i < len(nums) && nums[i] <= 0; i++ {
		if i-1 == 0 || nums[i] != nums[i-1] {
			twoSum(nums, i-1)
		}
	}
	return output
}

func twoSum(nums []int, i int) {
	val1 := nums[i]
	target := (0 + val1) * -1
	var left int
	var sum int
	right := len(nums) - 1

	for left < right {
		sum = nums[left] + nums[right]
		if sum == target {
			output = append(output, []int{val1, nums[left], nums[right]})
			left++
			right--
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
}
