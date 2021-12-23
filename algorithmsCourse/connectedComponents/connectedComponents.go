package connectedComponents

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputFilePath = "/Users/alexandrutodea/go/src/leetcode-algorithms/algorithmsCourse/connectedComponents/graph_week_1_prog_as.txt"

type Graph struct {
	vertices          []*Vertex
	verticesPositions []int
	sccs              []int
}

type Vertex struct {
	visited  bool
	incoming []int
	outgoing []int
}

func (g *Graph) getPositions() {
	for currentVertexIndex := len(g.vertices) - 1; currentVertexIndex > 0; currentVertexIndex-- {
		if g.vertices[currentVertexIndex].visited {
			continue
		}
		g.dfsForFindingVertexPosition(currentVertexIndex)
	}
}

func (g *Graph) dfsForFindingVertexPosition(currentVertexIndex int) {
	g.vertices[currentVertexIndex].visited = true
	for _, incomingVertexIndex := range g.vertices[currentVertexIndex].incoming {
		if g.vertices[incomingVertexIndex].visited {
			continue
		}
		g.dfsForFindingVertexPosition(incomingVertexIndex)
	}

	g.verticesPositions = append(g.verticesPositions, currentVertexIndex)
}

func (g *Graph) dfsForFindingSCCs(currentVertexIndex int) int {
	g.vertices[currentVertexIndex].visited = true
	numberVerticesInCluster := 1
	for _, outgoingVertex := range g.vertices[currentVertexIndex].outgoing {
		if g.vertices[outgoingVertex].visited {
			continue
		}
		numberVerticesInCluster += g.dfsForFindingSCCs(outgoingVertex)
	}

	return numberVerticesInCluster
}

func main() []int {
	data, err := parseFile(inputFilePath)
	graph := generateGraph(data, 875715) // for actual input
	// graph := generateGraph(data, 10) // for test2 file
	// graph := generateGraph(data, 9) // for test1 and test3 and test5 files
	// graph := generateGraph(data, 13) // for test4
	fmt.Print(graph)

	if err != nil {
		panic("error parsing the file")
	}

	graph.getPositions()
	graph.resetVisited()
	graph.getConnectedComponents()
	sort.Ints(graph.sccs)

	return graph.sccs
}

func (g *Graph) getConnectedComponents() {
	for i := len(g.verticesPositions) - 1; i >= 0; i-- {
		vertexIndex := g.verticesPositions[i]
		if g.vertices[vertexIndex].visited {
			continue
		}
		g.sccs = append(g.sccs, g.dfsForFindingSCCs(vertexIndex))
	}
}

func (g *Graph) resetVisited() {
	for i := 1; i < len(g.vertices); i++ {
		g.vertices[i].visited = false
	}
}

func generateGraph(input [][]int, numberVertices int) Graph {
	graph := Graph{vertices: make([]*Vertex, numberVertices), verticesPositions: []int{}, sccs: []int{}}

	for _, subArr := range input {
		from := subArr[0]
		to := subArr[1]

		if graph.vertices[from] == nil {
			graph.vertices[from] = &Vertex{}
		}

		if graph.vertices[to] == nil {
			graph.vertices[to] = &Vertex{}
		}

		graph.vertices[from].outgoing = append(graph.vertices[from].outgoing, to)
		graph.vertices[to].incoming = append(graph.vertices[to].incoming, from)
	}

	return graph
}

func parseFile(filePath string) ([][]int, error) {
	fmt.Println("Parsing file...")
	var fileRows [][]int

	file, err := os.Open(filePath)

	if err != nil {
		return [][]int{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineAsString := scanner.Text()
		row := strings.Fields(lineAsString)
		rowNew := make([]int, len(row))

		for i, valAsString := range row {
			valAsInt, err := strconv.Atoi(valAsString)

			if err != nil {
				return [][]int{}, err
			}

			rowNew[i] = valAsInt
		}

		fileRows = append(fileRows, rowNew)

		if err != nil {
			return [][]int{}, err
		}
	}

	if err := scanner.Err(); err != nil {
		return [][]int{}, err
	}
	fmt.Println("File parsed succesfully")
	return fileRows, nil
}
