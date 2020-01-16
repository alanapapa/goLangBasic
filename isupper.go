package piscine

func IsUpper(str string) bool {
	for _, el := range str {
		if el < 'A' || el > 'Z' {
			return false
		}
	}
	return true
}
