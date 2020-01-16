package piscine

func Any(f func(string) bool, arr []string) bool {
	result := false
	for _, e := range arr {
		if f(e) == true {
			result = true
		}
	}
	return result
}
