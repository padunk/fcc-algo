package basic

import (
	"regexp"
	"strings"
)

// ConfirmEnding takes string and target and
// return boolean if target is in the end of string
func ConfirmEnding(str, target string) bool {
	re := regexp.MustCompile(target + "$")
	return re.MatchString(str)
}

// ConfirmEndingSuffix use strings.HasSuffix that work
// pretty much the same like Strings.endsWith() in JS
func ConfirmEndingSuffix(str, target string) bool {
	return strings.HasSuffix(str, target)
}
