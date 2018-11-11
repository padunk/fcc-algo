package basic

import (
	"math"
	"sort"
)

// LargestOfFour takes nested slice of int and
// pick the biggest number in each slice and then collect it
// into a single slice
func LargestOfFour(arr [][]int) []int {
	result := make([]int, len(arr[0]))
	for i, a := range arr {
		sort.Ints(a)
		result[i] = a[len(a)-1]
	}
	return result
}

/*
using math.Max packages
*/

// LargestOfFourMax using math.Max packages
// that compare two values and return the biggest
func LargestOfFourMax(arr [][]int) []int {
	result := make([]int, len(arr))
	for i, a := range arr {
		// set the temp variable to -Infinity / lowest possible integer
		temp := math.Inf(-1)
		for _, n := range a {
			if math.Max(temp, float64(n)) > temp {
				temp = float64(n)
			}
		}
		result[i] = int(temp)
	}
	return result
}
