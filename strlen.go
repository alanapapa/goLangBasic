package piscine

func StrLen(str string) int {
	count := 0
	test := []rune(str)
	for first := range test {
		if first >= 0 {
			count++
		}
	}
	return count
}
