package piscine

func ConcatParams(args []string) string {
	result := ""
	sep := "\n"
	count := 0
	for range args {
		count++
	}
	for i, el := range args {
		result += el
		if i != count-1 {
			result += sep
		}
	}
	return result
}
