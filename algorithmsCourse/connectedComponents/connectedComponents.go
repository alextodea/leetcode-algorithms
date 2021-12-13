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
	topFiveConnectedComponents []int
	vertices                   map[int]*Vertex
}

type Vertex struct {
	label    int
	visited  bool
	outgoing []int
}

func main() []int {
	data, err := parseFile(inputFilePath)

	if err != nil {
		panic("error parsing the file")
	}

	graph := Graph{topFiveConnectedComponents: make([]int, 5), vertices: map[int]*Vertex{}}
	graph.construct(data)
	return graph.findTopFiveConnectedComponents()
}

func (g *Graph) findTopFiveConnectedComponents() []int {
	output := []int{}
	for label, vertex := range g.vertices {
		if vertex.visited {
			continue
		}

		output = append(output, g.bfs(label, 0))
	}

	sort.Ints(output)
	return output[:5]
}

func (g *Graph) bfs(label int, nrComponents int) int {
	if g.vertices[label].visited || len(g.vertices[label].outgoing) == 0 {
		return nrComponents
	}

	g.vertices[label].visited = true
	totalComponents := nrComponents + 1

	for _, label := range g.vertices[label].outgoing {
		totalComponents += g.bfs(label, totalComponents)
	}

	return totalComponents
}

func (g *Graph) construct(inputData [][]int) {
	for _, vertices := range inputData {
		sourceVertexAddress := g.addVertex(vertices[0])
		destinationVertexAddress := g.addVertex(vertices[1])
		g.addEge(sourceVertexAddress, destinationVertexAddress)
	}
}

func (g *Graph) addVertex(label int) int {
	_, vertexExists := g.vertices[label]

	if !vertexExists {
		g.vertices[label] = &Vertex{label: label, outgoing: []int{}, visited: false}
	}

	return label
}

func (g *Graph) addEge(sourceLabel int, destinationLabel int) {
	if sourceLabel == destinationLabel {
		return
	}

	g.vertices[sourceLabel].outgoing = append(g.vertices[sourceLabel].outgoing, destinationLabel)
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
