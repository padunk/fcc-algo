package inter

// SumAll take 1 argument and return a function or
// take 2 arguments and sum it all
func SumAll(s []int) int {
	if s[0] > s[1] {
		s[0], s[1] = s[1], s[0]
	}
	total := 0
	for i := s[0]; i <= s[1]; i++ {
		total += i
	}
	return total
}
