package basic

import "strings"

// Mutation check if arr[0] strings contains all of the letters
// in arr[1] element, ignoring the case
func Mutation(arr []string) bool {
	arr[0] = strings.ToLower(arr[0])
	arrToCheck := strings.Split(strings.ToLower(arr[1]), "")
	for _, s := range arrToCheck {
		if !strings.Contains(arr[0], s) {
			return false
		}
	}
	return true
}
