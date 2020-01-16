package piscine

func ToLower(s string) string {
	result := ""
	for _, str := range s {
		if str >= 'A' && str <= 'Z' {
			str += 32
		}
		result += string(rune(str))
	}
	return result
}
