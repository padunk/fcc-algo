package basic

import "strings"

// RepeatStringNumTimes takes string and an integer
// return repeating string as many as 'num' times
func RepeatStringNumTimes(str string, num int) string {
	var result string
	if num <= 0 {
		return ""
	}
	for i := 0; i < num; i++ {
		result += str
	}
	return result
}

// RepeatStringNumTimes2 using strings.Repeat method
// that worked same like Strings.repeat() in JS
func RepeatStringNumTimes2(str string, num int) string {
	return strings.Repeat(str, num)
}
