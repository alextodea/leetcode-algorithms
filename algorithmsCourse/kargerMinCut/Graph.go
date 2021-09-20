package kargermincut

import "fmt"

type Input struct {
	nrEdges    int
	nrVertices int
	fileRows   [][]int
}

type Edge struct {
	from int
	to   int
}

// type Vertex struct {
// 	label int
// }

type Graph struct {
	edges    []Edge
	vertices []int
}

func (g *Graph) initialize(input Input) {
	g.addVertices(input.nrVertices)
	g.addEdges(input.fileRows)
	fmt.Println("s")
}

func (g *Graph) addEdges(rows [][]int) {
	var edgeIndex int
	for _, rowArray := range rows {
		for i := 1; i < len(rowArray); i++ {
			g.edges[edgeIndex].from = rowArray[0]
			g.edges[edgeIndex].to = rowArray[i]
			edgeIndex++
		}
	}
}

func (g *Graph) addVertices(nrVertices int) {
	for i := 0; i < nrVertices; i++ {
		g.vertices[i] = i
	}
}


// initialize all vertices in 'vertices' array with -1 and decrease each time a child is added
// use path compression in order to reduce the amortize time complexity needed in order to access parent from child
func (g *Graph) mergeVertices() {
	// for _, edge := range g.edges {

	// }
}

// func (g *Graph) minCuts() {
// 	nrvertices := g.vertices[len(g.vertices)-1] // last element in vertices array will always equal total nr vertices

// 	for nrvertices > 2 {

// 		nrvertices--
// 	}
// }
