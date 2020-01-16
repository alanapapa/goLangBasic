package piscine

func AppendRange(min, max int) []int {
	var result []int
	var nil []int
	if min < max {
		for i := min; i < max; i++ {
			result = append(result, i)
		}
	} else if min >= max {
		return nil
	}
	return result
}
