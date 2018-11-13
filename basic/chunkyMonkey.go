package basic

import "math"

/*
-First we have to create nested slice with length of 'chunkLength',
because in Go we can't push element into 'nil' slice except with append
-The first if checked if size equal or less than 0, return empty slice
-The second if checked if size bigger or equal length of arr, return arr
-Loop through arr and increment it with 'size' step
-Push the sliced arr into result
*/

// ChunkArrayInGroups will splits 'arr' into groups of slice
// with 'size' length each
func ChunkArrayInGroups(arr []int, size int) [][]int {
	if size <= 0 {
		return [][]int{}
	}
	chunkLength := int(math.Ceil(float64(len(arr)) / float64(size)))
	result := make([][]int, chunkLength)

	if size >= len(arr) {
		result[0] = arr
		return result
	}
	i := 0
	for j := 0; j < len(arr); j += size {
		if j+size >= len(arr) {
			result[i] = arr[j:]
		} else {
			result[i] = arr[j : j+size]
		}
		i++
	}
	return result
}

/*
Solution with append

	if size <= 0 {
		return [][]int{}
	}
	// No need to intialized the length of result
	var result [][]int

	if len(arr) <= size {
		// create an empty slice inside
		result = append(result, []int{})
		result[0] = append(result[0], arr...)
		return result
	}
	i := 0
	for j := 0; j < len(arr); j += size {
		// create an empty slice inside
		result = append(result, []int{})
		if j+size >= len(arr) {
			result[i] = append(result[i], arr[j:]...)
		} else {
			result[i] = append(result[i], arr[j : j+size]...)
		}
		i++
	}
	return result
*/
