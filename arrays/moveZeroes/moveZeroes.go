package arrays

// func moveZeroes(nums []int) []int {

// 	if nums == nil || len(nums) == 0 {
// 		return nil
// 	}

// 	if len(nums) == 1 {
// 		return nums
// 	}

// 	for i := len(nums) - 1; i > 0; i-- {
// 		if nums[i-1] == 0 {
// 			for j := i; j < len(nums) && nums[j] != 0; j++ {
// 				temp := nums[j]
// 				nums[j] = nums[j-1]
// 				nums[j-1] = temp
// 			}
// 		}
// 	}
// 	return nums
// }
