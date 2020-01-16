package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	okbro := []rune(str)
	for first, element := range okbro {
		if first >= 0 {
			z01.PrintRune(element)
		}
	}
}
