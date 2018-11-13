package basic

// ReverseString takes string and
// return reverse string pattern
func ReverseString(str string) string {
	var result []byte
	// convert string to byte, we can also use
	// strings.Split to split strings into substring
	b := []byte(str)
	for i := len(b) - 1; i >= 0; i-- {
		result = append(result, b[i])
	}
	// convert byte to string again
	return string(result)
}
