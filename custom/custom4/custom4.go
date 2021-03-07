package custom

import "sort"

func howManySwaps(arr []int32) int64 {
	var nrMinimumSwaps int
	elementsMap := map[int32]int{}
	sortedArr := make([]int32, len(arr))
	copy(sortedArr, arr)
	sort.Sort(SortByInt32(sortedArr))

	for i, val := range arr {
		elementsMap[val] = i
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != sortedArr[i] {
			nrMinimumSwaps++
			temp := arr[i]
			arr[i] = arr[elementsMap[int32(i)]]
			arr[elementsMap[int32(i)]] = temp
			elementsMap[arr[i]] = elementsMap[sortedArr[i]]
			elementsMap[sortedArr[i]] = i
		}
	}
	return int64(nrMinimumSwaps)
}

type SortByInt32 []int32

func (a SortByInt32) Len() int           { return len(a) }
func (a SortByInt32) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByInt32) Less(i, j int) bool { return a[i] < a[j] }
