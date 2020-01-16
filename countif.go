package piscine

func CountIf(f func(string) bool, tab []string) int {
	result := 0
	for _, e := range tab {
		if f(e) == true {
			result++
		}
	}
	return result
}
