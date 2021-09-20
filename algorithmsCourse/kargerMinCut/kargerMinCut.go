package kargermincut

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var filePath = "data.txt"

func main() {
	input, err := parseFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	graph := Graph{edges: make([]Edge, input.nrEdges), vertices: make([]int, input.nrVertices)}

	graph.initialize(input)
	graph.mergeVertices()
	// fmt.Println(graph.minCuts())
}

func parseFile(filePath string) (Input, error) {

	var nrEdges int
	nrVertices := 1
	var fileRows [][]int

	file, err := os.Open(filePath)

	if err != nil {
		return Input{}, err
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
				return Input{}, err
			}

			rowNew[i] = valAsInt
		}

		fileRows = append(fileRows, rowNew)

		nrEdges += len(row) - 1
		nrVertices++

		if err != nil {
			return Input{}, err
		}
	}

	if err := scanner.Err(); err != nil {
		return Input{}, err
	}

	return Input{nrEdges, nrVertices, fileRows}, nil
}

func (g *Graph) getMinCuts() int {
	return -1
}
