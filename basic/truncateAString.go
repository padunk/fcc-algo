package basic

import "strings"

func TruncateString(str string, num int) string {
	if num > len(str) {
		return str
	}
	return strings.Join(strings.Split(str, "")[num:], "") + "..."
}
