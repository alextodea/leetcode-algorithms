package custom

func Solution(S string) int {
	if len(S) < 3 {
		return -1
	}

	mapOfDigrams := make(map[string]int)
	maximumDistance := -1
	var currentDigram string

	for i := 1; i < len(S); i++ {
		currentDigram = S[i-1 : i+1]
		if _, digramExistsInMap := mapOfDigrams[currentDigram]; !digramExistsInMap {
			mapOfDigrams[currentDigram] = i - 1
		} else {
			thisDistance := (i - 1) - mapOfDigrams[currentDigram]
			if thisDistance > maximumDistance {
				maximumDistance = thisDistance
			}
		}
	}

	return maximumDistance
}
