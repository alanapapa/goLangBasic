package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, e := range *ptr {
		if e != "" {
			count++
		}
	}
	mass := make([]string, count)
	i := 0
	for _, e := range *ptr {
		if e != "" {
			mass[i] = e
			i++
		}
	}

	*ptr = mass
	return count
}
