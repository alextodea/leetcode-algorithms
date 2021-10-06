package main

import (
	"fmt"
	kargermincut "leetcodeAlgorithms/algorithmsCourse/kargerMinCut"
	"sort"
)

var testFile = "kargerMinCut"

func main() {
	minCuts := kargermincut.Run(testFile, 10)
	sort.Ints(minCuts)
	fmt.Println(minCuts)
}
