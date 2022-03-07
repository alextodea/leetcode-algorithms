package djikstra

import (
	"bufio"
	"container/heap"
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
	from  int
	to    int
	score int
	index int // index of edge in priority queue
}

func (edges *Edges) addEdgesFromBatch(from, indexInEdgesArray int, batch []string) error {
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
		from:  from,
		to:    to,
		score: score,
	}

	(*edges)[indexInEdgesArray] = &edge // add edge at location 'indexInEdgesArray' within edges arr. this is happening because we pre-assigned memory for this array

	nextIndex := indexInEdgesArray + 1
	batchTail := batch[1:]

	return edges.addEdgesFromBatch(from, nextIndex, batchTail)
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
		sourceVertexLabel := len(g.adjacencyList)
		err := thisVertexEdges.addEdgesFromBatch(sourceVertexLabel, 0, rowArrayOfStrings[1:]) // skip first element because that's the vertex

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

func (g *Graph) getShortestPathDistance(from, destinationVertex int) int {
	path := make([]int, len(g.adjacencyList))
	costs := make([]int, len(path))
	pq := make(PriorityQueue, len(*g.adjacencyList[from]))

	// add edges from source vertex
	for i, edge := range *g.adjacencyList[from] {
		pq[i] = edge
		pq[i].index = i
	}

	for len(pq) > 0 {
		// if we reached our target vertex, return cost
		if path[destinationVertex] != 0 {
			return costs[destinationVertex]
		}
		heap.Init(&pq)
		edge := heap.Pop(&pq).(*Edge)

		// if source vertex exists at path index and cost at given index is smaller than that from edge then continue to next iteration
		if path[edge.to] != 0 {
			if edge.score >= costs[edge.to] {
				continue
			}
		}

		path[edge.to] = edge.from
		costs[edge.to] = edge.score

		// add edges from source vertex
		for _, newEdge := range *g.adjacencyList[edge.to] {
			if path[newEdge.from] == newEdge.to {
				continue
			}
			newEdge.score = newEdge.score + edge.score
			heap.Push(&pq, newEdge)
		}

	}

	fmt.Println("path: ", path)
	fmt.Println("costs: ", costs)
	return costs[destinationVertex]
}

func main(datasetName string, destinationVertex int) int {
	fileName := "/" + datasetName + ".txt"
	pwd, _ := os.Getwd()
	filePath := pwd + fileName
	graph := initGraph()

	err := graph.createGraphFromFile(filePath)

	if err != nil {
		panic("error creating graph from file")
	}

	return graph.getShortestPathDistance(1, destinationVertex)

}

// implement priority queue

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the edge that contains the lowest distance between two vertices.
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	edge := x.(*Edge)
	edge.index = n
	*pq = append(*pq, edge)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	edge := old[n-1]
	old[n-1] = nil  // avoid memory leak
	edge.index = -1 // for safety
	*pq = old[0 : n-1]
	return edge
}
