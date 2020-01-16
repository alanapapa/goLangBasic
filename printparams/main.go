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

	for i := 1; i < b; i++ {
		for _, el := range a[i] {
			z01.PrintRune(el)
		}

		z01.PrintRune('\n')
	}
}
