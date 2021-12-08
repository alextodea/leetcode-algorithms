package arrays

import "sort"

func medianSlidingWindow(nums []int, k int) []float64 {
	input := Input{nums: nums, output: []float64{}, windowSize: k, window: make([]int, k)}
	return input.getMedianArrayOfEachWindow(0, k-1)
}

type Input struct {
	nums       []int
	output     []float64
	window     []int
	windowSize int
}

func getWindowMedian(window []int) float64 {
	if len(window)%2 == 0 {
		return float64(window[len(window)/2-1]+window[len(window)/2]) / 2
	}
	return float64(window[len(window)/2])
}

func (i *Input) getMedianArrayOfEachWindow(left, right int) []float64 {
	if right >= len(i.nums) {
		return i.output
	}

	for j, val := range i.nums[left : right+1] {
		i.window[j] = val
	}

	sort.Ints(i.window)
	windowMedian := getWindowMedian(i.window)

	i.output = append(i.output, windowMedian)

	return i.getMedianArrayOfEachWindow(left+1, right+1)
}
