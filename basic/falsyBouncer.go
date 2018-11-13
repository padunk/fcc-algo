package basic

import (
	"reflect"
)

type object struct {
	I   int
	F   float64
	B   bool
	S   string
	Isl []int
	Ssl []string
}

// Bouncer takes structs with multiple kind of data
// and remove the 'falsy' value from it
func Bouncer(obj object) []interface{} {
	v := reflect.ValueOf(obj)
	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		if !isZeroOfUnderlyingType(v.Field(i).Interface()) {
			values[i] = v.Field(i).Interface()
		}
	}
	return values
}

// this function is from SO
// https://stackoverflow.com/questions/13901819/quick-way-to-detect-empty-values-via-reflection-in-go
func isZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
