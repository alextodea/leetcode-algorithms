package arrays

import "sort"

func threeSum(nums []int) [][]int {
	var output [][]int
	if len(nums) == 0 {
		return output
	}

	sort.Ints(nums)

	twoSums := map[int][]int{}

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val == nums[i-1] || val == 0 {
			continue
		}

		diff := 0 - val
		inverseOfDiff := diff * -1
		if twoSumsArr, exists := twoSums[inverseOfDiff]; exists {
			output = append(output, []int{i, twoSumsArr[0], twoSumsArr[1]})
		} else {
			
		}
	}
	return output
}
