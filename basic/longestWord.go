package basic

import (
	"sort"
	"strings"
)

/*
package strings.Split will split string into slice of substrings
work the same like Strings.split in JS
slice is like Array in JS
*/

// FindLongestWordLength takes the string and
// return the length of the longest word
func FindLongestWordLength(str string) int {
	result := 0
	strSlice := strings.Split(str, " ")
	for _, s := range strSlice {
		if len(s) > result {
			result = len(s)
		}
	}
	return result
}

/*
with sort functions in Go
code from https://gobyexample.com/sorting-by-functions
*/

type strSlice []string

func (s strSlice) Len() int {
	return len(s)
}
func (s strSlice) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
func (s strSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// FindLongestWordSort will sort the slice of substrings
// based on its length
func FindLongestWordSort(str string) int {
	s := strings.Split(str, " ")
	sort.Sort(strSlice(s))
	return len(s[0])
}
