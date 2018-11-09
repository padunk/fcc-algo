package basic

// ConvertToF takes celsius and
// return Fahrenheit
func ConvertToF(c float64) float64 {
	return 9.0/5.0*c + 32
}

/* integer need to used decimal point otherwise
it will truncated the zero. 9/5 will become 1.
*/
