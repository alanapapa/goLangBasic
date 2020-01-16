package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	a := os.Args
	b := 0

	for range a {
		b++
	}

	for i := b; i > 1; i-- {
		for _, el := range a[i-1] {
			z01.PrintRune(el)
		}

		z01.PrintRune('\n')
	}
}
