package djikstra

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	adjacencyList []*Edges
}

type Edges []*Edge

type Edge struct {
	to    int
	score int
}

func (edges *Edges) addEdgesFromBatch(indexInEdgesArray int, batch []string) error {
	if len(batch) == 0 {
		return nil
	}

	batchHead := batch[0]
	splittedStrings := strings.Split(batchHead, ",")
	to, err := strconv.Atoi(splittedStrings[0])

	if err != nil {
		fmt.Print("Cannot convert " + splittedStrings[0] + " to int")
		return err
	}

	score, err := strconv.Atoi(splittedStrings[1])

	if err != nil {
		fmt.Print("Cannot convert " + splittedStrings[1] + " to int")
		return err
	}

	edge := Edge{
		to:    to,
		score: score,
	}

	(*edges)[indexInEdgesArray] = &edge // add edge at location 'indexInEdgesArray' within edges arr. this is happening because we pre-assigned memory for this array

	nextIndex := indexInEdgesArray + 1
	batchTail := batch[1:]

	return edges.addEdgesFromBatch(nextIndex, batchTail)
}

func initGraph() Graph {
	return Graph{adjacencyList: []*Edges{}}
}

func (g *Graph) createGraphFromFile(filePath string) error {

	fmt.Println("Parsing file...")
	g.adjacencyList = append(g.adjacencyList, &Edges{}) // append dummy vertex at index 0 in order to maintain correct naming

	file, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineAsString := scanner.Text()
		rowArrayOfStrings := strings.Fields(lineAsString)
		thisVertexEdges := make(Edges, len(rowArrayOfStrings[1:]))
		err := thisVertexEdges.addEdgesFromBatch(0, rowArrayOfStrings[1:]) // skip first element because that's the vertex

		if err != nil {
			return err
		}

		g.adjacencyList = append(g.adjacencyList, &thisVertexEdges)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("Graph created succesfully")
	return nil
}

func main() {
	fileName := "/datasetMain.txt"
	pwd, _ := os.Getwd()
	filePath := pwd + fileName
	graph := initGraph()

	err := graph.createGraphFromFile(filePath)

	if err != nil {
		panic("error creating graph from file")
	}
}
