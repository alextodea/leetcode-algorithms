package custom

func Solution(A []int) int {
	if len(A) <= 1 {
		return len(A)
	}

	var maxNrTurbulences int
	var nrTurbulencesForCurrentPeriod int
	var prevHeightVolatility int
	var currentHeightVolatility int

	for i := 1; i < len(A); i++ {
		currentHeightVolatility = determineHeightVolatility(A[i-1], A[i])
		if currentHeightVolatility == 0 || prevHeightVolatility == currentHeightVolatility {
			if nrTurbulencesForCurrentPeriod > maxNrTurbulences {
				maxNrTurbulences = nrTurbulencesForCurrentPeriod
			}
			nrTurbulencesForCurrentPeriod = 1
		}
		prevHeightVolatility = currentHeightVolatility
		nrTurbulencesForCurrentPeriod++
	}

	return maxNrTurbulences
}

func determineHeightVolatility(h1, h2 int) int {
	if h1 > h2 {
		return 1
	} else if h2 > h1 {
		return -1
	}
	return 0
}
