package hashMaps

func repeatedNTimes(A []int) int {
	mapOfUniqueValues := map[int]int{}

	for _, val := range A {
		_, valExists := mapOfUniqueValues[val]
		if valExists {
			return val
		} else {
			mapOfUniqueValues[val] = 1
		}
	}
	return 0
}
