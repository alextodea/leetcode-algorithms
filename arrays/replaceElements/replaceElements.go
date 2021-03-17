package arrays

// Given an array arr, replace every element in that array with the greatest element among the elements to its right,
// and replace the last element with -1.
// After doing so, return the array.
func replaceElements(arr []int) []int {
	if arr == nil {
		return nil
	}

	if len(arr) == 1 {
		arr[0] = -1
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[i] = arr[j]
			}
		}
	}

	arr[len(arr)-1] = -1

	return arr
}
