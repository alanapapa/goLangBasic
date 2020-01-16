package piscine

func FirstRune(s string) rune {
	var a rune
	for _, i := range s {
		a = i
		break
	}
	return a
}
