package main

import (
	"fmt"

	"github.com/padunk/fcc-algo/basic"
	"github.com/padunk/fcc-algo/inter"
)

func main() {
	fmt.Println(basic.ConvertToF(100))
	fmt.Println(inter.Destroyer([]int{1, 2, 3, 1, 2, 3}, 2, 3))
}
