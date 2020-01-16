package piscine

func Join(strs []string, sep string) string {
	result := ""
	count := 0
	for range strs {
		count++
	}
	for first, el := range strs {
		result += el
		if first != count-1 {
			result += sep
		}
	}
	return result
}
