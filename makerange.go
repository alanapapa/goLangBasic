package piscine

func MakeRange(min, max int) []int {
	var nil []int
	if min < max {
		result := make([]int, max-min)
		count := 0
		for i := min; i < max; i++ {
			result[count] = i
			count++
			nil = result
		}
	} else if min < 0 || min >= max {
		return nil
	}
	return nil
}
