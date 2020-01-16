package piscine

func IsAlpha(str string) bool {
	for _, el := range str {
		if el < '0' || el > '9' && el < 'A' || el > 'Z' && el < 'a' || el > 'z' {
			return false
		}
	}
	return true
}
