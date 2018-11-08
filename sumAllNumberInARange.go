package main

import "fmt"

func main() {
	fmt.Println(sumAll([]int{1, 4}))
	fmt.Println(sumAll([]int{4, 1}))
	fmt.Println(sumAll([]int{5, 10}))
	fmt.Println(sumAll([]int{10, 5}))
}

func sumAll(s []int) int {
	if s[0] > s[1] {
		s[0], s[1] = s[1], s[0]
	}
	total := 0
	for i := s[0]; i <= s[1]; i++ {
		total += i
	}
	return total
}
