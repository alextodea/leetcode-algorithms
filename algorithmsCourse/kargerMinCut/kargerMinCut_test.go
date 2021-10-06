package kargermincut

import (
	"fmt"
	"sort"
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

	graph := Graph{edges: make([]Edge, input.nrEdges), vertices: make([]Subset, input.nrVertices)}

	graph.initialize(input)
}

func TestKargerMinCutWithTestFileNr1(t *testing.T) {
	minCuts := Run("test1", 10)
	sort.Ints(minCuts)
	if minCuts[0] != 2 {
		t.Error("result of test1 should be 2")
	}
}

func TestKargerMinCutWithTestFileNr2(t *testing.T) {
	minCuts := Run("test2", 10)
	sort.Ints(minCuts)
	if minCuts[0] != 1 {
		t.Error("result of test2 should be 1")
	}
}

func TestKargerMinCutWithTestFileNr3(t *testing.T) {
	minCuts := Run("test3", 10)
	sort.Ints(minCuts)
	if minCuts[0] != 3 {
		t.Error("result of test3 should be 3")
	}
}

func TestKargerMinCutWithTestFileNr4(t *testing.T) {
	minCuts := Run("test4", 10)
	sort.Ints(minCuts)
	if minCuts[0] != 2 {
		t.Error("result of test4 should be 2")
	}
}
