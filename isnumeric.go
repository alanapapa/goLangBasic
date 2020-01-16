package piscine

func IsNumeric(str string) bool {
	for _, el := range str {
		if el < '0' || el > '9' {
			return false
		}
	}
	return true
}
