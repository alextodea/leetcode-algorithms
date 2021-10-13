package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) < 1 || numCourses == 0 {
		return false
	}

	adjacencyList := AdjacencyList(make(AdjacencyList, numCourses))
	adjacencyList.initialize(numCourses, prerequisites)
	adjacencyList.determinePrerequisites()
	return false // temporary
}
