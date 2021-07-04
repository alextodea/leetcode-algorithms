package countinverses

import (
	"fmt"
	"testing"
)

func TestGetArrayOfIntegersFromFile(t *testing.T) {
	var integersArray, integersArrayError = getArrayOfIntegersFromFile("data.txt")

	if integersArrayError != nil {
		t.Error("there was an error when testing getArrayOfIntegersFromFile: /n", integersArrayError)
	}

	expectedResult := []int{54044, 14108, 79294, 29649, 25260, 60660, 2995, 53777, 49689, 9083, 16122, 90436, 4615, 40660}

	for i, val := range integersArray {
		if val != expectedResult[i] {
			t.Error("array values are different than expected")
		}
	}
}

func TestCountNumberOfInversions(t *testing.T) {
	var integersArray, integersArrayError = getArrayOfIntegersFromFile("data.txt")
	if integersArrayError != nil {
		t.Error("there was an error when testing getArrayOfIntegersFromFile: /n", integersArrayError)
	}

	_, nrInversions := countNumberOfInversions(integersArray)
	fmt.Println(nrInversions)
	expectedResult := 3

	if expectedResult != nrInversions {
		t.Error("res should be 3")
	}
}

func TestMain(t *testing.T) {
	main()
}
