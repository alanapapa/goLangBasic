package piscine

func IsLower(str string) bool {
	for _, el := range str {
		if el < 'a' || el > 'z' {
			return false
		}
	}
	return true
}
