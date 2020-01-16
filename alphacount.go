package piscine

func AlphaCount(str string) int {
	letters := []rune(str)
	counter := 0
	for _, element := range letters {
		if element >= 65 && element <= 90 || element >= 97 && element <= 122 {
			counter++
		}
	}
	return counter
}
