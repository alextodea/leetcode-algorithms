package kargermincut

import (
	"fmt"
	"testing"
)

var inputDataFilePath = "./kargerMinCut.txt"

func TestParseFileFunc(t *testing.T) {
	input, err := parseFile(inputDataFilePath)

	if err != nil {
		t.Error("Parse file error")
	}

	fmt.Print(input)
}

func TestParseInitializeGraph(t *testing.T) {
	input, err := parseFile(inputDataFilePath)

	if err != nil {
		t.Error("Parse file error")
	}

	graph := Graph{edges: make([]Edge, input.nrEdges), vertices: make([]int, input.nrVertices)}
	
	graph.initialize(input)
}