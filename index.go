package piscine

func Lenght(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}
func Index(s string, toFind string) int {
	result := -1
	lenS := Lenght(s)
	lenF := Lenght(toFind)
	arrS := []rune(s)
	arrF := []rune(toFind)
	if s == "" {
		return 0
	}
	for index, letter := range s {
		if letter == arrF[0] && index+lenF <= lenS {
			j := 0
			count := 0
			for i := index; i < lenF; i++ {
				if arrS[i] == arrF[j] {
					count++
				}
				j++
			}
			if count == j {
				result = index
				break
			}
		}
	}
	return result
}
