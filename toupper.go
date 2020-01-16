package piscine

func ToUpper(s string) string {
	result := ""
	for _, str := range s {
		if str >= 'a' && str <= 'z' {
			str = str - 32
		}
		result += string(rune(str))
	}
	return result
}
