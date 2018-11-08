# FreeCodeCamp Intermediate Algorithm Challenges in Go Language
This is my journye to learn how Go works, maybe not perfect or maybe a bit naive, but through this process I push my self to learn much of Go packages and tricks.

Comments, edit, whatever you want to text me are welcome!
Pull the request now :-)

## How it works
```sh
go get github.com/padunk/fcc-intermediate
```

In your new `main.go` files
```go
package main

import (
   "fmt"

   "GOPATH/padunk/fcc-intermediate"
)

func main() {
   fmt.Println(inter.SumAll([]int{1, 4}))
}
```

You can get all the test from FreeCodeCamp [website](https://learn.freecodecamp.org) or copy paste below test.

### Sum All Number in a Range
```go
   fmt.Println(inter.SumAll([]int{1, 4}))
	fmt.Println(inter.SumAll([]int{4, 1}))
	fmt.Println(inter.SumAll([]int{5, 10}))
	fmt.Println(inter.SumAll([]int{10, 5}))
```

### Diff Two Arrays
```go
   fmt.Println(inter.DiffIntArray([]int{1, 2, 3, 5}, []int{1, 2, 3, 4, 5}))
	fmt.Println(inter.DiffStringArray([]string{"diorite", "andesite", "grass", "dirt", "pink wool", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(inter.DiffStringArray([]string{"andesite", "grass", "dirt", "pink wool", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(inter.DiffStringArray([]string{"diorite", "andesite", "grass", "dirt", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(inter.DiffStringArray([]string{}, []string{"snuffleupagus", "cookie monster", "elmo"}))
	fmt.Println(inter.DiffIntArray([]int{1, 3}, []int{7}))
```

### Seek and Destroy
```go
   fmt.Println(inter.Destroyer([]int{1, 2, 3, 1, 2, 3}, 2, 3))
	fmt.Println(inter.Destroyer([]int{1, 2, 3, 5, 1, 2, 3}, 2, 3))
	fmt.Println(inter.Destroyer([]int{3, 5, 1, 2, 2}, 2, 3, 5))
	fmt.Println(inter.Destroyer([]int{2, 3, 2, 3}, 2, 3))
```
