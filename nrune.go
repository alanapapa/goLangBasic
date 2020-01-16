package piscine

func NRune(s string, n int) rune {
	strRune := []rune(s)
	counter := 0
	for range s {
		counter++
	}
	var help rune
	if n <= 0 {
		return 0
	} else if counter >= n {
		return strRune[n-1]
	}
	return help
}
