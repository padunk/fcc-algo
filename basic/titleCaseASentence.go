package basic

import "strings"

/*
there is strings.Title in Go but have bug see https://golang.org/pkg/strings/#Title
- So we lowercase the string with strings.ToLower()
- split it into slice (like array in JS)
- loop it with range
- use strings.ToTitle on s[0], but since s[0] is byte
- we convert it to string first with string(s[byte index])
- concat it with s[i:] same like .slice / .substring in JS
- and join it
*/

// TitleCase takes a string and
// return title case string
func TitleCase(str string) string {
	result := strings.Split(strings.ToLower(str), " ")
	for i, s := range result {
		result[i] = strings.ToTitle(string(s[0])) + s[1:]
	}
	return strings.Join(result, " ")
}
