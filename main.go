package main

import (
	"fmt"

	"github.com/padunk/fcc_algo_in_golang/inter"
)

func main() {
	fmt.Println(inter.Destroyer([]int{1, 2, 3, 1, 2, 3}, 2, 3))
	fmt.Println(inter.Destroyer([]int{1, 2, 3, 5, 1, 2, 3}, 2, 3))
	fmt.Println(inter.Destroyer([]int{3, 5, 1, 2, 2}, 2, 3, 5))
	fmt.Println(inter.Destroyer([]int{2, 3, 2, 3}, 2, 3))
}
