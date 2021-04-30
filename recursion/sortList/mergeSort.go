package recursion

// time complexity O(N log N)
// space complexity O()
func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	half := len(nums) / 2

	firstArr := sortArray(nums[:half])
	secondArr := sortArray(nums[half:])

	return merge(firstArr, secondArr)

}

func merge(firstArr, secondArr []int) []int {
	mergedListLength := len(firstArr) + len(secondArr)
	mergedList := make([]int, mergedListLength)
	var firstArrPivot int
	var secondArrPivot int
	var mergedListPivot int

	for firstArrPivot < len(firstArr) && secondArrPivot < len(secondArr) {
		if firstArr[firstArrPivot] < secondArr[secondArrPivot] {
			mergedList[mergedListPivot] = firstArr[firstArrPivot]
			firstArrPivot++
		} else {
			mergedList[mergedListPivot] = secondArr[secondArrPivot]
			secondArrPivot++
		}
		mergedListPivot++
	}

	for firstArrPivot < len(firstArr) {
		mergedList[mergedListPivot] = firstArr[firstArrPivot]
		firstArrPivot++
		mergedListPivot++
	}

	for secondArrPivot < len(secondArr) {
		mergedList[mergedListPivot] = secondArr[secondArrPivot]
		secondArrPivot++
		mergedListPivot++
	}

	return mergedList
}
