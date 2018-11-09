# FreeCodeCamp Intermediate Algorithm Challenges in Go Language
This is my journey to learn how Go works, maybe not perfect or maybe a bit naive, but through this process I push my self to learn much of Go packages, tricks and how to think in Go.

Comments, edit or whatever you want to text me are welcome!
Pull the request now :-)

## How it works
```sh
cd $GOPATH/src
git clone https://github.com/padunk/fcc-algo
cd fcc-algo
code main.go
```

Open `main.go` file with your favorite text editor.<br>
It will show:
```go
package main

import (
   "fmt"
   // change the GOPATH to your gopath directory
   "GOPATH/fcc-algo/basic"
   "GOPATH/fcc-algo/inter"
)

func main() {
        fmt.Println(basic.ConvertToF(100))
	fmt.Println(inter.Destroyer([]int{1, 2, 3, 1, 2, 3}, 2, 3))
}
```

Run `main.go` in you cli
```sh
go run main.go
```

You can get all the test from FreeCodeCamp [website](https://learn.freecodecamp.org) or copy paste below test.

## Basic Algorithm Scripting

## Convert Celsius to Fahrenheit
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/convert-celsius-to-fahrenheit)
```go
   fmt.Println(basic.ConvertToF(-30))
   fmt.Println(basic.ConvertToF(-10))
   fmt.Println(basic.ConvertToF(0))
   fmt.Println(basic.ConvertToF(20))
   fmt.Println(basic.ConvertToF(30))
```

## Reverse a String
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/reverse-a-string)
```go
   fmt.Println(basic.ReverseString("hello"))
   fmt.Println(basic.ReverseString("Howdy"))
   fmt.Println(basic.ReverseString("Greetings from Earth"))
```

## Factorialize a Number
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/factorialize-a-number)
```go
   fmt.Println(basic.Factorialize(5))
   fmt.Println(basic.Factorialize(10))
   fmt.Println(basic.Factorialize(20))
   fmt.Println(basic.Factorialize(0))
```


## Intermediate Algorithm Scripting

### Sum All Number in a Range
[link]()
```go
        fmt.Println(inter.SumAll([]int{1, 4}))
	fmt.Println(inter.SumAll([]int{4, 1}))
	fmt.Println(inter.SumAll([]int{5, 10}))
	fmt.Println(inter.SumAll([]int{10, 5}))
```

### Diff Two Arrays
[link]()
```go
   	fmt.Println(inter.DiffIntArray([]int{1, 2, 3, 5}, []int{1, 2, 3, 4, 5}))
	fmt.Println(inter.DiffStringArray([]string{"diorite", "andesite", "grass", "dirt", "pink wool", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(inter.DiffStringArray([]string{"andesite", "grass", "dirt", "pink wool", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(inter.DiffStringArray([]string{"diorite", "andesite", "grass", "dirt", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(inter.DiffStringArray([]string{}, []string{"snuffleupagus", "cookie monster", "elmo"}))
	fmt.Println(inter.DiffIntArray([]int{1, 3}, []int{7}))
```

### Seek and Destroy
[link]()
```go
   	fmt.Println(inter.Destroyer([]int{1, 2, 3, 1, 2, 3}, 2, 3))
	fmt.Println(inter.Destroyer([]int{1, 2, 3, 5, 1, 2, 3}, 2, 3))
	fmt.Println(inter.Destroyer([]int{3, 5, 1, 2, 2}, 2, 3, 5))
	fmt.Println(inter.Destroyer([]int{2, 3, 2, 3}, 2, 3))
```
