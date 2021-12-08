package arrays

import "testing"

func Test1(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := medianSlidingWindow(nums, k)
	expected := []float64{1.00000, -1.00000, -1.00000, 3.00000, 5.00000, 6.00000}

	for i, val := range result {
		if val != expected[i] {
			t.Error("values should be the same")
		}
	}
}

func Test2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 2, 3, 1, 4, 2}
	k := 3
	result := medianSlidingWindow(nums, k)
	expected := []float64{2.00000, 3.00000, 3.00000, 3.00000, 2.00000, 3.00000, 2.00000}

	for i, val := range result {
		if val != expected[i] {
			t.Error("values should be the same")
		}
	}
}

func Test3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 2, 3, 1, 4, 2}
	k := 4
	result := medianSlidingWindow(nums, k)
	expected := []float64{2.00000, 3.00000, 3.00000, 3.00000, 2.00000, 3.00000, 2.00000}

	for i, val := range result {
		if val != expected[i] {
			t.Error("values should be the same")
		}
	}
}
