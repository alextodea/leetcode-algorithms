package arrays

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	numsMap := make(map[int]int, len(nums))

	for index, value := range nums {
		complement := target - value
		if complementValueFromMap, exists := numsMap[complement]; exists {
			return []int{index, complementValueFromMap}
		}
		numsMap[value] = index
	}

	return []int{-1, -1}
}
