package basic

/*
this code using type switch
more explanation in https://tour.golang.org/methods/16
*/

// BooWho accept any interface and
// return true if 'i' is type Boolean
func BooWho(i interface{}) bool {
	switch i.(type) {
	case bool:
		return true
	default:
		return false
	}
}
