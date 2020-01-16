package piscine

func TrimAtoi(s string) int {
	res := 0
	help := 1
	for _, i := range s {
		if i >= 48 && i <= 57 {
			res = (int(i) - 48) + res*10
		} else if i == '-' && res == 0 {
			help = -1
		}
	}
	return res * help
}
