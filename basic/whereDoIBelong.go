package basic

import "sort"

// GetIndexToIns takes slice of integer and
// return index where 'num' is belong to the slice after
// sorting incrementally
func GetIndexToIns(arr []int, num int) int {
	sort.Ints(arr)
	if len(arr) == 0 || arr[0] >= num {
		return 0
	} else if arr[len(arr)-1] < num {
		return len(arr)
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == num || arr[i] > num {
			return i
		}
	}
	return 0
}
