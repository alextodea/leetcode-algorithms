package hashMaps

func numIdenticalPairs(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	pairsMappings := addPairsToMappings(nums)
	var nrIdenticalPairs int

	for _, val := range pairsMappings {
		nrIdenticalPairs += val
	}

	return nrIdenticalPairs
}

func addPairsToMappings(nums []int) map[int]int {
	pairsMappings := map[int]int{}

	for i := 0; i < len(nums)-1; i++ {
		_, valExistsInMappings := pairsMappings[nums[i]]
		if !valExistsInMappings {
			pairsMappings[nums[i]] = 0
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				pairsMappings[nums[i]]++
			}
		}
	}
	return pairsMappings
}
