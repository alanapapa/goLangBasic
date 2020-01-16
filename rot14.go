package piscine

func Rot14(str string) string {

	len := 0
	result := ""

	for range str {
		len++
	}

	for _, e := range str {
		if e >= 97 && e <= 108 || e >= 65 && e <= 76 {
			result += string(e + 14)
		}
		if e >= 109 && e <= 122 || e >= 77 && e <= 90 {
			result += string(e - 12)
		}
		if e == 32 {
			result += string(32)
		}
		if e >= 48 && e <= 57 {
			result += string(e)
		}
	}
	return result
}
