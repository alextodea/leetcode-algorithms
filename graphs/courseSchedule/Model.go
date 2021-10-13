package courseschedule

type AdjacencyList []Vertex

type Vertex struct {
	parent        int
	order         int
	outgoingEdges []*Vertex
}

func (al *AdjacencyList) initialize(numCourses int, prerequisites [][]int) {
	al.initializeVertices(numCourses)
	al.initializeEdges(prerequisites)
}

func (al *AdjacencyList) initializeEdges(prerequisites [][]int) {
	for _, prerequisite := range prerequisites {
		srcVertexIndex := prerequisite[1]
		dstVertexIndex := prerequisite[0]
		if (*al)[srcVertexIndex].outgoingEdges == nil {
			(*al)[srcVertexIndex].outgoingEdges = []*Vertex{}
		}
		(*al)[srcVertexIndex].outgoingEdges = append((*al)[srcVertexIndex].outgoingEdges, &(*al)[dstVertexIndex])
	}
}

func (al *AdjacencyList) initializeVertices(numCourses int) {
	for i := 0; i < len(*al); i++ {
		(*al)[i] = Vertex{parent: i}
	}
}

func (al *AdjacencyList) dfsTopo(vertex Vertex) {

}

func (al *AdjacencyList) determinePrerequisites() {
	for _, vertex := range *al {
		al.dfsTopo(vertex)
	}
}
