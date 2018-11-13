package basic

import "strings"

// TruncateString takes string and cut a string according to 'num'
// if num bigger than string return string
// else return truncate string with '...'
func TruncateString(str string, num int) string {
	if num > len(str) {
		return str
	}
	return strings.Join(strings.Split(str, "")[num:], "") + "..."
}
