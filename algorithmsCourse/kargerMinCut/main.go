package kargermincut

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run(testFileName string, seed int) []int {

	var filePath = "/Users/alexhome/go/src/leetcodeAlgorithms/algorithmsCourse/kargerMinCut/datasets/" + testFileName + ".txt"
	input, err := parseFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	minCuts := []int{}

	for i := 0; i < seed; i++ {
		graph := Graph{edges: make([]Edge, input.nrEdges), vertices: make([]Subset, input.nrVertices)}
		graph.initialize(input)
		minCuts = append(minCuts, graph.kargerMinCut())
	}

	return minCuts
}

func parseFile(filePath string) (Input, error) {
	fmt.Println("Parsing file...")
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
	fmt.Println("File parsed succesfully")
	return Input{nrEdges, nrVertices, fileRows}, nil
}
