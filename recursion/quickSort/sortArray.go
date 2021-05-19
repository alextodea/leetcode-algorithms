package recursion

import "math/rand"

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	less, equal, greater := []int{}, []int{}, []int{}

	pivot := nums[rand.Intn(len(nums)-1)]

	for i := 0; i < len(nums); i++ {
		arrVal := nums[i]

		if arrVal < pivot {
			less = append(less, arrVal)
		} else if arrVal > pivot {
			greater = append(greater, arrVal)
		} else {
			equal = append(equal, arrVal)
		}
	}

	return append(append(sortArray(less), equal...), sortArray(greater)...)
}
