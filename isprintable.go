package piscine

func IsPrintable(str string) bool {
	for _, el := range str {
		if el < 32 {
			return false
		}
	}
	return true
}
