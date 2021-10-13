package courseschedule

import "testing"

func Test1(t *testing.T) {
	numCourses := 2
	prerequisites := make([][]int, 1)
	prerequisites[0] = []int{1, 0}
	result := canFinish(numCourses, prerequisites)
	expectedResult := true

	if result != expectedResult {
		t.Error("result should be", expectedResult)
	}
}

func Test2(t *testing.T) {
	numCourses := 2
	prerequisites := make([][]int, 2)
	prerequisites[0] = []int{1, 0}
	prerequisites[1] = []int{0, 1}
	result := canFinish(numCourses, prerequisites)
	expectedResult := true

	if result != expectedResult {
		t.Error("result should be", expectedResult)
	}
}


