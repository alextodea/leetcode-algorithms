package arrays

type Nums struct {
	list   []int
	target int
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	numsConstructor := &Nums{list: nums, target: target}
	firstIndex := numsConstructor.getIndex(true, 0, len(nums)-1)
	secondIndex := numsConstructor.getIndex(false, 0, len(nums)-1)

	return []int{firstIndex, secondIndex}
}

func (n *Nums) getIndex(isFirstIndex bool, lowerBound, upperBound int) int {
	if lowerBound > upperBound {
		return -1
	}

	middleIndex := lowerBound + ((upperBound - lowerBound) / 2)
	middleIndexValue := n.list[middleIndex]

	if middleIndexValue == n.target {
		if isFirstIndex == true {
			if lowerBound == middleIndex || middleIndexValue != n.list[middleIndex-1] {
				return middleIndex
			}
			upperBound = middleIndex - 1
		} else {
			if upperBound == middleIndex || middleIndexValue != n.list[middleIndex+1] {
				return middleIndex
			}
			lowerBound = middleIndex + 1
		}
		return n.getIndex(isFirstIndex, lowerBound, upperBound)
	} else if middleIndexValue < n.target {
		return n.getIndex(isFirstIndex, middleIndex+1, upperBound)
	}

	return n.getIndex(isFirstIndex, lowerBound, middleIndex-1)
}
