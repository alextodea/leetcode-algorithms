package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	g := Graph{adjacencyList: make([]Vertex, numCourses)}
	g.initialize(numCourses, prerequisites)

	for i := 0; i < len(g.adjacencyList); i++ {
		if g.hasCycle(i) {
			return false
		}
	}

	return true
}
