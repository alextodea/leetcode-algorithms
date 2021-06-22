package countinverses

import "testing"

func TestGetArrayOfIntegersFromFile(t *testing.T) {
	result, err := getArrayOfIntegersFromFile("data-test.txt")

	if err != nil {
		t.Error("func returned an error")
	}

	expectedResult := []int{54044, 14108, 79294, 29649, 25260, 60660, 2995, 53777, 49689, 9083, 16122, 90436, 4615, 40660}

	for val, i := range result {
		if val != expectedResult[i] {
			t.Error("array values are different than expected")
		}
	}
}
