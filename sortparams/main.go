package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	count := 0
	arg := os.Args[1:]
	for index := range arg {
		if index >= 0 {
			count++
		}
	}
	for i := 1; i <= count; i++ {
		out := os.Args[i]
		for a := 0; a < count; a++ {
			for b := 0; b < count; b++ {
				if arg[a] < arg[b] {
					arg[a], arg[b] = arg[b], arg[a]
				}
			}
		}
		for _, j := range out {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
