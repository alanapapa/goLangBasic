package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	count := 0
	slice := []rune(s)
	for index := range slice {
		if index >= 0 {
			count++
		}
	}
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0
		}
		n = n*10 + int(ch)
	}
	if s0[0] == '-' {
		n = -n
	}
	return n
}
