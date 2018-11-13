package basic

import "fmt"

// FrankenSplice will takes two arrays and combine them
// It takes arr2 as the base array and arr1 as a copied array
// combine happen at index 'n' in arr2
func FrankenSplice(arr1, arr2 []int, n int) []int {
	var result []int
	if n > 0 {
		result = append(result, arr2[:n]...)
		result = append(result, arr1...)
		result = append(result, arr2[n:]...)
	} else if n == 0 {
		result = append(result, arr1...)
		result = append(result, arr2...)
	}
	fmt.Printf("arr1: %v, arr2: %v\n", arr1, arr2)
	return result
}
