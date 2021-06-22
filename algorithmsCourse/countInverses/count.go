package countinverses

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var filePath = "data.txt"

func main() {
	arr, err := getArrayOfIntegersFromFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(arr)
}

func getArrayOfIntegersFromFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	outputArr := make([]int, 0)

	for scanner.Scan() {
		processedLine, err := processLine(scanner.Text())

		if err != nil {
			return []int{}, err
		}

		outputArr = append(outputArr, processedLine)
	}

	if err := scanner.Err(); err != nil {
		return []int{}, err
	}

	fmt.Println("succesfully created array of integers from file")
	return outputArr, nil
}

func processLine(line string) (int, error) {
	val, err := strconv.Atoi(line)
	if err != nil {
		return -1, err
	}
	return val, nil
}
