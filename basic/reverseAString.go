package basic

// ReverseString takes string and
// return reverse string pattern
func ReverseString(str string) string {
	var resultByte []byte
	b := []byte(str)
	for i := len(b) - 1; i >= 0; i-- {
		resultByte = append(resultByte, b[i])
	}

	return string(resultByte)
}
