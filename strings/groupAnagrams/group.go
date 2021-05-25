package strings

import "sort"

func groupAnagrams(strs []string) [][]string {
	mapOfGrouppedAnagrams := make(map[string][]string)

	for _, anagramString := range strs {
		anagramStringAsArrayOfBytes := []byte(anagramString)
		sort.Slice(anagramStringAsArrayOfBytes, func(i, j int) bool {
			return anagramStringAsArrayOfBytes[i] < anagramStringAsArrayOfBytes[j]
		})

		sortedAnagramString := string(anagramStringAsArrayOfBytes)
		mapOfGrouppedAnagrams[sortedAnagramString] = append(mapOfGrouppedAnagrams[sortedAnagramString], anagramString)
	}

	output := make([][]string, 0, len(mapOfGrouppedAnagrams))

	for _, grouppedAnagramsArray := range mapOfGrouppedAnagrams {
		output = append(output, grouppedAnagramsArray)
	}

	return output
}
