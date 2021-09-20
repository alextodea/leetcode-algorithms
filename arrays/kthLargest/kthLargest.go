package arrays

import "math/rand"

type Input struct {
	nums        []int
	targetIndex int
}

// based on rSelect algorithm which runs in O(n) time and space complexity is O(n)
func findKthLargest(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || k > len(nums) {
		return -1
	}

	if len(nums) == 1 {
		return 1
	}

	input := Input{nums, len(nums) - k}

	return input.rSelect(0, len(nums)-1)
}

// Fails when min=0 and max=1, find alternative for random function
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (input *Input) rSelect(left, right int) int {
	randomIndex := randInt(left, right)
	input.swap(left, randomIndex)
	limit := input.partition(left, right)

	if input.targetIndex == limit {
		return input.nums[input.targetIndex]
	} else if input.targetIndex < limit {
		return input.rSelect(left, limit)
	} else {
		return input.rSelect(limit+1, right)
	}
}

func (input *Input) partition(left, right int) int {
	p := left
	i := left + 1
	j := left + 1

	for j <= right {
		if input.nums[j] < input.nums[p] {
			input.swap(i, j)
			i++
		}
		j++
	}
	input.swap(i-1, p)
	return i - 1
}

func (input *Input) swap(x, y int) {
	temp := input.nums[x]
	input.nums[x] = input.nums[y]
	input.nums[y] = temp
}
