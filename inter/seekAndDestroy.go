package inter

// Destroyer takes array and filter out the 'remove' in the array
// return array with element that is not 'remove'
func Destroyer(arr []int, remove ...int) []int {
	var rem []int
	for index := 0; index < len(remove); index++ {
		rem = append(rem, remove[index])
	}
	for _, r := range rem {
		arr = filter(arr, r, removeCB)
	}
	return arr
}

type callback func(x, y int) bool

func filter(s []int, r int, fn callback) []int {
	var result []int
	for _, value := range s {
		if fn(value, r) {
			result = append(result, value)
		}
	}
	return result
}

func removeCB(i, r int) bool {
	if i != r {
		return true
	}
	return false
}
