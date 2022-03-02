package djikstra

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type Edge struct {
// 	to    int
// 	score int
// }

// type Vertex []*Edge

// type Graph struct {
// 	vertices []*Vertex
// }

// func (g *Graph) init() {}

func main() {
	fileName := "/datasetMain.txt"
	pwd, _ := os.Getwd()
	filePath := pwd + fileName
	data, err := parseFile(filePath)

	if err != nil {
		panic("error parsing the file")
	}

	fmt.Print(data)
}

// todo error parsing file
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
