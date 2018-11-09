package basic

/*
The output will be uint64 because the result will
be very big numbers if we want to find 20!
*/

// Factorialize takes i as integer
// and return factorial of 'i'
func Factorialize(i int) uint64 {
	var result uint64 = 1
	arr := make([]int, i)
	for idx := 1; idx <= i; idx++ {
		arr[idx-1] = idx
	}
	for _, a := range arr {
		result *= uint64(a)
	}
	return result
}

// Recursive function from https://gobyexample.com/recursion

// RecurFactorial same as factorialize
// but call the function recursively.
func RecurFactorial(i int) uint64 {
	if i == 0 {
		return 1
	}
	return uint64(i) * RecurFactorial(i-1)
}
