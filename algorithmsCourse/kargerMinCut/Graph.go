package kargermincut

import (
	"fmt"
	"math/rand"
	"time"
)

type Input struct {
	nrEdges    int
	nrVertices int
	fileRows   [][]int
}

type Edge struct {
	from int
	to   int
}

type Subset struct {
	parent int
	rank   int
}

type Graph struct {
	edges    []Edge
	vertices []Subset
}

func (g *Graph) initialize(input Input) {
	fmt.Println("Initializing graph....")
	g.addVertices(input.nrVertices)
	g.addEdges(input.fileRows)
	fmt.Println("Graph initialized succesfully")
}

func (g *Graph) addEdges(rows [][]int) {
	var edgeIndex int
	var row []int
	for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
		row = rows[rowIndex]
		for colIndex := 1; colIndex < len(row); colIndex++ {
			g.edges[edgeIndex].from = row[0]
			g.edges[edgeIndex].to = row[colIndex]
			edgeIndex++
		}
	}
}

func (g *Graph) addVertices(nrVertices int) {
	for i := 0; i < nrVertices; i++ {
		g.vertices[i] = Subset{parent: i}
	}
}

func (g *Graph) kargerMinCut() int {
	fmt.Println("Running karger min cut algorithm...")
	nrVertices := len(g.vertices) - 1
	for nrVertices > 2 {
		time.Sleep(time.Millisecond * 21)
		rand.Seed(time.Now().Unix())
		i := rand.Intn(len(g.edges))
		s1ParentVertexId := g.findSubsetParentVertexId(g.edges[i].from)
		s2ParentVertexId := g.findSubsetParentVertexId(g.edges[i].to)
		g.popEdgeOut(i)

		if !g.areVerticesPartOfSameSubset(s1ParentVertexId, s2ParentVertexId) {
			g.unionOfSubsets(s1ParentVertexId, s2ParentVertexId)
			nrVertices--
		}
	}

	var minCuts int

	fmt.Println("Identifying nr min cuts...")
	countedEdges := map[int]bool{}
	for i := 0; i < len(g.edges); i++ {
		sourceVertexIndex := g.edges[i].from
		destVertexIndex := g.edges[i].to
		subset1 := g.findSubsetParentVertexId(sourceVertexIndex)
		subset2 := g.findSubsetParentVertexId(destVertexIndex)
		edgeExists := countedEdges[sourceVertexIndex]
		if subset1 != subset2 && !edgeExists {
			countedEdges[destVertexIndex] = true
			minCuts++
		}
	}

	fmt.Println("karger min algorithm ran succesfully")
	return minCuts
}

func (g *Graph) areVerticesPartOfSameSubset(s1ParentVertexId, s2ParentVertexId int) bool {
	return s1ParentVertexId == s2ParentVertexId
}

func (g *Graph) unionOfSubsets(s1ParentVertexId, s2ParentVertexId int) {
	if g.vertices[s1ParentVertexId].rank > g.vertices[s2ParentVertexId].rank {
		g.vertices[s2ParentVertexId].parent = g.vertices[s1ParentVertexId].parent
	} else if g.vertices[s1ParentVertexId].rank < g.vertices[s2ParentVertexId].rank {
		g.vertices[s1ParentVertexId].parent = g.vertices[s2ParentVertexId].parent
	} else {
		g.vertices[s1ParentVertexId].parent = g.vertices[s2ParentVertexId].parent
		g.vertices[s2ParentVertexId].rank++
	}
}

func (g *Graph) findSubsetParentVertexId(vertexIndex int) int {
	if g.vertices[vertexIndex].parent != vertexIndex {
		g.vertices[vertexIndex].parent = g.findSubsetParentVertexId(g.vertices[vertexIndex].parent)
	}

	return g.vertices[vertexIndex].parent
}

func (g *Graph) popEdgeOut(edgeIndex int) {
	edge := g.edges[edgeIndex]
	lenEdges := len(g.edges) - 1
	g.edges[edgeIndex] = g.edges[lenEdges]
	g.edges[lenEdges] = edge
	g.edges = g.edges[:lenEdges]
}
