package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	result := os.Args[0]
	for _, el := range result {
		z01.PrintRune(el)
	}
	z01.PrintRune('\n')
}
