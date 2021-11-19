package arrays

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var maxArea int
	var i int
	j := len(height) - 1

	for i < j {
		thisArea := (j - i) * getSmallest(height[i], height[j])
		if thisArea > maxArea {
			maxArea = thisArea
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}

func getSmallest(x1, x2 int) int {
	if x1 < x2 {
		return x1
	}
	return x2
}
