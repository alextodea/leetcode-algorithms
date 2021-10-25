package courseschedule

type Graph struct {
	adjacencyList []Vertex
}

type Vertex struct {
	visited       bool
	outgoingEdges []int
}

func (g *Graph) initialize(numCourses int, prerequisites [][]int) {
	g.initializeEdges(prerequisites)
}

func (g *Graph) initializeEdges(prerequisites [][]int) {
	for _, prerequisite := range prerequisites {
		srcVertexIndex := prerequisite[0]
		dstVertexIndex := prerequisite[1]
		if g.adjacencyList[srcVertexIndex].outgoingEdges == nil {
			g.adjacencyList[srcVertexIndex].outgoingEdges = []int{}
		}
		g.adjacencyList[srcVertexIndex].outgoingEdges = append(g.adjacencyList[srcVertexIndex].outgoingEdges, dstVertexIndex)
	}
}

func (g *Graph) initializeVertices(numCourses int) {
	for i := 0; i < len(g.adjacencyList); i++ {
		g.adjacencyList[i] = Vertex{visited: false, outgoingEdges: []int{}}
	}
}

func (g *Graph) hasCycle(i int) bool {
	if g.adjacencyList[i].visited {
		return true
	}

	if len(g.adjacencyList[i].outgoingEdges) == 0 {
		return false
	}

	g.adjacencyList[i].visited = true

	for _, prerequisiteVertexIndex := range g.adjacencyList[i].outgoingEdges {
		if g.hasCycle(prerequisiteVertexIndex) {
			return true
		}
	}

	g.adjacencyList[i].visited = false
	g.adjacencyList[i].outgoingEdges = []int{}

	return false
}
