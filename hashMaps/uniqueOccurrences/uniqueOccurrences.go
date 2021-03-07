package hashMaps

type Occurences map[int]int

func uniqueOccurrences(arr []int) bool {
	if len(arr) < 2 {
		return true
	}

	occurrencesTable := addOccurencesToTable(arr)

	return areOccurencesUnique(occurrencesTable)
}

func areOccurencesUnique(occurrencesTable Occurences) bool {
	uniqueOccurencesAarr := []int{}
	for _, val := range occurrencesTable {
		if arrayContainsValue(uniqueOccurencesAarr, val) {
			return false
		} else {
			uniqueOccurencesAarr = append(uniqueOccurencesAarr, val)
		}
	}
	return true
}

func arrayContainsValue(arr []int, val int) bool {
	if len(arr) < 1 {
		return false
	}

	for _, arrVal := range arr {
		if arrVal == val {
			return true
		}
	}
	return false
}

// func areOccurencesUnique(occurrencesTable Occurences) bool {
// 	uniqueOccurencesAarr := []int{}

// 	for _, val := range occurrencesTable {
// 		uniqueOccurencesAarr = append(uniqueOccurencesAarr, val)
// 	}

// 	sort.Ints(uniqueOccurencesAarr)

// 	for i := 1; i < len(uniqueOccurencesAarr); i++ {
// 		prev := uniqueOccurencesAarr[i-1]
// 		curr := uniqueOccurencesAarr[i]

// 		if prev == curr {
// 			return false
// 		}
// 	}

// 	return true
// }

func addOccurencesToTable(arr []int) Occurences {
	occurencesTable := make(map[int]int)

	for _, arrVal := range arr {
		addOccurenceToTable(arrVal, occurencesTable)
	}

	return occurencesTable
}

func addOccurenceToTable(arrVal int, occurencesTable Occurences) {
	if _, exists := occurencesTable[arrVal]; exists {
		occurencesTable[arrVal]++
	} else {
		occurencesTable[arrVal] = 1
	}
}
