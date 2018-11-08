package inter

// DiffIntArray compare two slices and return the different element
// parameter is INTEGER slices
func DiffIntArray(arr1, arr2 []int) []int {
	var result []int
	if len(arr1) == 0 {
		return arr2
	} else if len(arr2) == 0 {
		return arr1
	}
	diff := func(a1, a2 []int) {
		for index := 0; index < len(a1); index++ {
			if intIndexOf(a2, a1[index]) == -1 {
				result = append(result, a1[index])
			}
		}
	}
	diff(arr1, arr2)
	diff(arr2, arr1)

	return result
}

// DiffStringArray compare two slices and return the different element
// parameter is STRING slices
func DiffStringArray(arr1, arr2 []string) []string {
	var result []string
	if len(arr1) == 0 {
		result = arr2
		return result
	} else if len(arr2) == 0 {
		result = arr1
		return result
	}
	diff := func(a1, a2 []string) {
		for index := 0; index < len(a1); index++ {
			if stringIndexOf(a2, a1[index]) == -1 {
				result = append(result, a1[index])
			}
		}
	}
	diff(arr1, arr2)
	diff(arr2, arr1)
	return result
}

func intIndexOf(slice []int, num int) int {
	indexAt := 0
	for i := 0; i < len(slice); i++ {
		if num == slice[i] {
			indexAt = i
			return indexAt
		}
	}
	indexAt = -1
	return indexAt
}

func stringIndexOf(slice []string, str string) int {
	indexAt := 0
	for i := 0; i < len(slice); i++ {
		if str == slice[i] {
			indexAt = i
			return indexAt
		}
	}
	indexAt = -1
	return indexAt
}
