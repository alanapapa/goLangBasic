package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printS(str string) {
	arrayStr := []rune(str)
	count := 0
	for range arrayStr {
		count++
	}

	for i := 0; i < count; i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	yes := true
	no := false
	if nbr%2 == 0 {
		return no
	} else {
		return yes
	}
}

func main() {
	args := os.Args
	lengthOfArg := 0
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	for range args {
		lengthOfArg++
	}
	if isEven(lengthOfArg) == true {
		printS(EvenMsg)
	} else {
		printS(OddMsg)
	}
}
