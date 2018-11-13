package basic

// FindElement takes slice of arr and a f function
// return first element that satisfy f function
func FindElement(arr []int, f func(int) bool) int {
	for _, a := range arr {
		if f(a) {
			return a
		}
	}
	return 0
}
