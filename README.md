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

### Convert Celsius to Fahrenheit
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/convert-celsius-to-fahrenheit)
```go
   fmt.Println(basic.ConvertToF(-30))
   fmt.Println(basic.ConvertToF(-10))
   fmt.Println(basic.ConvertToF(0))
   fmt.Println(basic.ConvertToF(20))
   fmt.Println(basic.ConvertToF(30))
```

### Reverse a String
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/reverse-a-string)
```go
   fmt.Println(basic.ReverseString("hello"))
   fmt.Println(basic.ReverseString("Howdy"))
   fmt.Println(basic.ReverseString("Greetings from Earth"))
```

### Factorialize a Number
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/factorialize-a-number)
```go
   fmt.Println(basic.Factorialize(5))
   fmt.Println(basic.Factorialize(10))
   fmt.Println(basic.Factorialize(20))
   fmt.Println(basic.Factorialize(0))
```

### Find the Longest Word in a String
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/find-the-longest-word-in-a-string)
```go
   fmt.Println(basic.FindLongestWordLength("The quick brown fox jumped over the lazy dog"))
   fmt.Println(FindLongestWordLength("May the force be with you"))
   fmt.Println(FindLongestWordLength("Google do a barrel roll"))
   fmt.Println(FindLongestWordSort("What is the average airspeed velocity of an unladen swallow"))
   fmt.Println(FindLongestWordSort("What if we try a super-long word such as otorhinolaryngology"))
```

### Return Largest Numbers in Arrays
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/return-largest-numbers-in-arrays)
```go
   fmt.Println(basic.LargestOfFour([[13, 27, 18, 26], [4, 5, 1, 3], [32, 35, 37, 39], [1000, 1001, 857, 1]])) // [27, 5, 39, 1001]
   fmt.Println(basic.LargestOfFour([[4, 9, 1, 3], [13, 35, 18, 26], [32, 35, 97, 39], [1000000, 1001, 857, 1]])) // [9, 35, 97, 1000000]
   fmt.Println(basic.LargestOfFour([[17, 23, 25, 12], [25, 7, 34, 48], [4, -10, 18, 21], [-72, -3, -17, -10]])) // [25, 48, 21, -3]
```

### Confirm the Ending
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/confirm-the-ending)
```go
fmt.Println(basic.ConfirmEnding("Bastian", "n")) // true
fmt.Println(basic.ConfirmEnding("Congratulation", "on")) // true
fmt.Println(basic.ConfirmEnding("Connor", "n")// false
fmt.Println(basic.ConfirmEnding("Walking on water and developing software from a specification are easy if both are frozen", "specification")) // false
fmt.Println(basic.ConfirmEnding("He has to give me a new name", "name")) // true
fmt.Println(basic.ConfirmEnding("Open sesame", "same")) // true
fmt.Println(basic.ConfirmEnding("Open sesame", "pen")) // false
fmt.Println(basic.ConfirmEnding("Open sesame", "game")) // false
fmt.Println(basic.ConfirmEnding("If you want to save our world, you must hurry. We dont know how much longer we can withstand the nothing", "mountain")) // false
fmt.Println(basic.ConfirmEnding("Abstraction", "action")) // true
```

### Repeat a String
[link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/repeat-a-string-repeat-a-string)
```go
   fmt.Println(basic.RepeatStringNumTimes("*", 3)) // "***"
   fmt.Println(basic.RepeatStringNumTimes("abc", 3)) // "abcabcabc"
   fmt.Println(basic.RepeatStringNumTimes("abc", 4)) // "abcabcabcabc"
   fmt.Println(basic.RepeatStringNumTimes("abc", 1)) // "abc"
   fmt.Println(basic.RepeatStringNumTimes("*", 8)) // "********"
   fmt.Println(basic.RepeatStringNumTimes("abc", -2)) // ""
```

### Truncate a String
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/truncate-a-string)
[code](https://github.com/padunk/fcc-algo/blob/master/basic/truncateAString.go)
```go
   fmt.Println(basic.TruncateString("A-tisket a-tasket A green and yellow basket", 8)) // "A-tisket..."
   fmt.Println(basic.TruncateString("Peter Piper picked a peck of pickled peppers", 11)) // "Peter Piper..."
   fmt.Println(basic.TruncateString("A-tisket a-tasket A green and yellow basket", len("A-tisket a-tasket A green and yellow basket"))) //  "A-tisket a-tasket A green and yellow basket"
   fmt.Println(basic.TruncateString("A-tisket a-tasket A green and yellow basket", len("A-tisket a-tasket A green and yellow basket") + 2)) // "A-tisket a-tasket A green and yellow basket"
   fmt.Println(basic.TruncateString("A-", 1)) // "A..."
   fmt.Println(basic.TruncateString("Absolutely Longer", 2)) // "Ab..."
```

### Finders Keepers
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/finders-keepers)
```go
   fmt.Println(basic.FindElement([]int{1, 3, 5, 8, 9}, func(n int) bool { return n%2 == 0 })) // 8
```

### Boo Who
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/boo-who)
```go
   fmt.Println(basic.BooWho(false)) // true
   fmt.Println(basic.BooWho(true)) // true
   fmt.Println(basic.BooWho([]int{1, 2, 3})) // false
   fmt.Println(basic.BooWho(map[string]int{ "a": 1 })) // false
   fmt.Println(basic.BooWho(1)) // false
   fmt.Println(basic.BooWho("a")) // false
   fmt.Println(basic.BooWho("true")) // false
   fmt.Println(basic.BooWho("false"))// false
```

### Title Case a Sentence
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/title-case-a-sentence)
```go
   fmt.Println(TitleCase("I'm a little tea pot"))
	fmt.Println(TitleCase("sHoRt AnD sToUt"))
   fmt.Println(TitleCase("HERE IS MY HANDLE HERE IS MY SPOUT"))
```

### Slice and Splice
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/slice-and-splice)
```go
   fmt.Println(basic.FrankenSplice([]int{1, 2, 3}, []int{4, 5}, 1)) //[4, 1, 2, 3, 5]
   fmt.Println(basic.FrankenSplice([]int{1, 2}, []int{4, 5}, 2)) // [4, 5, 1, 2]
```

### Falsy Bouncer
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/falsy-bouncer)
For this challenges, I don't think it can be done in Go, so I try to replicate
it as best as I can do.
This 'falsy' values (0, "", empty slice) will become 'nil'.
But please don't do this in Go, this is now how Golang really works.
```go
type object struct {
	I   int
	F   float64
	B   bool
	S   string
	Isl []int
	Ssl []string
}

func main() {
   x := object{9, 1.5, true, "", []int{}, []string{"hey", "world"}}
   y := object{0, 1.5, false, "golang", []int{1, 2}, []string{}}
   z := object{0, 0, true, "golang", []int{1, 2}, []string{"hey", "world"}}
   
   fmt.Println(basic.Bouncer(x)) // [9, 1.5, true, <nil>, [], ["hey", "world"]]
   fmt.Println(basic.Bouncer(y)) // [<nil>, 1.5, <nil>, "golang", [1, 2], []]
   fmt.Println(basic.Bouncer(z)) // [<nil>, <nil>, true, "golang", [1, 2], ["hey", "world"]]
}

```

### Where Do I Belong
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/where-do-i-belong)
```go
   fmt.Println(basic.GetIndexToIns([]int{10, 20, 30, 40, 50}, 35)) // 3
   fmt.Println(basic.GetIndexToIns([]int{10, 20, 30, 40, 50}, 30)) // 2
   fmt.Println(basic.GetIndexToIns([]int{40, 60}, 50)) // 1
   fmt.Println(basic.GetIndexToIns([]int{3, 10, 5}, 3)) // 0
   fmt.Println(basic.GetIndexToIns([]int{5, 3, 20, 3}, 5)) // 2
   fmt.Println(basic.GetIndexToIns([]int{2, 20, 10}, 19)) // 2
   fmt.Println(basic.GetIndexToIns([]int{2, 5, 10}, 15)) // 3
   fmt.Println(basic.GetIndexToIns([]int{}, 1)) // 0
```

### Mutations
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/mutations)
```go
   fmt.Println(basic.Mutation([]string{"hello", "hey"})) //false
   fmt.Println(basic.Mutation([]string{"hello", "Hello"})) // true
   fmt.Println(basic.Mutation([]string{"zyxwvutsrqponmlkjihgfedcba", "qrstu"}))// true
   fmt.Println(basic.Mutation([]string{"Mary", "Army"})) // true
   fmt.Println(basic.Mutation([]string{"Mary", "Aarmy"})) // true
   fmt.Println(basic.Mutation([]string{"Alien", "line"})) // true
   fmt.Println(basic.Mutation([]string{"floor", "for"})) // true
   fmt.Println(basic.Mutation([]string{"hello", "neo"})) //false
   fmt.Println(basic.Mutation([]string{"voodoo", "no"})) //false
```

### Chunky Monkey
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/basic-algorithm-scripting/chunky-monkey)
```go
   fmt.Println(basic.ChunkArrayInGroups([]int{0, 1, 2, 3, 4, 5}, 3)) // [[0, 1, 2], [3, 4, 5]]
   fmt.Println(basic.ChunkArrayInGroups([]int{0, 1, 2, 3, 4, 5}, 2)) // [[0, 1], [2, 3], [4, 5]]
   fmt.Println(basic.ChunkArrayInGroups([]int{0, 1, 2, 3, 4, 5}, 4)) // [[0, 1, 2, 3], [4, 5]]
   fmt.Println(basic.ChunkArrayInGroups([]int{0, 1, 2, 3, 4, 5, 6}, 3)) // [[0, 1, 2], [3, 4, 5], [6]]
   fmt.Println(basic.ChunkArrayInGroups([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, 4)) // [[0, 1, 2, 3], [4, 5, 6, 7], [8]]
   fmt.Println(basic.ChunkArrayInGroups([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, 2)) // [[0, 1], [2, 3], [4, 5], [6, 7], [8]]
```

<br>
<br>
## Intermediate Algorithm Scripting

### Sum All Numbers in a Range
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/intermediate-algorithm-scripting/sum-all-numbers-in-a-range)
```go
        fmt.Println(inter.SumAll([]int{1, 4}))
	fmt.Println(inter.SumAll([]int{4, 1}))
	fmt.Println(inter.SumAll([]int{5, 10}))
	fmt.Println(inter.SumAll([]int{10, 5}))
```

### Diff Two Arrays
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/intermediate-algorithm-scripting/diff-two-arrays)
```go
   	fmt.Println(inter.DiffIntArray([]int{1, 2, 3, 5}, []int{1, 2, 3, 4, 5}))
	fmt.Println(inter.DiffStringArray([]string{"diorite", "andesite", "grass", "dirt", "pink wool", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(inter.DiffStringArray([]string{"andesite", "grass", "dirt", "pink wool", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(inter.DiffStringArray([]string{"diorite", "andesite", "grass", "dirt", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(inter.DiffStringArray([]string{}, []string{"snuffleupagus", "cookie monster", "elmo"}))
	fmt.Println(inter.DiffIntArray([]int{1, 3}, []int{7}))
```

### Seek and Destroy
[FCC link](https://learn.freecodecamp.org/javascript-algorithms-and-data-structures/intermediate-algorithm-scripting/seek-and-destroy)
```go
   	fmt.Println(inter.Destroyer([]int{1, 2, 3, 1, 2, 3}, 2, 3))
	fmt.Println(inter.Destroyer([]int{1, 2, 3, 5, 1, 2, 3}, 2, 3))
	fmt.Println(inter.Destroyer([]int{3, 5, 1, 2, 2}, 2, 3, 5))
	fmt.Println(inter.Destroyer([]int{2, 3, 2, 3}, 2, 3))
```
