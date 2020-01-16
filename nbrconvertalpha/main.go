package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args[1:]
	var constant int
	constant = 96

	for _, word := range name {
		result := 0
		for _, l := range word {
			if os.Args[1] == "--upper" {
				constant = 64
				name = os.Args[2:]
			}
			result = int(l-48) + result*10
		}
		if result > 26 {
			z01.PrintRune(' ')
		} else if result == 0 {

		} else {
			z01.PrintRune(rune(result + constant))
		}

	}
	z01.PrintRune('\n')
}
