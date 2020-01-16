package piscine

func Max(arr []int) int {
	mass := arr
	len := 0
	res := 0
	for range arr {
		len++
	}
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if mass[i] < mass[j] {
				mass[i], mass[j] = mass[j], mass[i]
			}
		}
	}

	for _, e := range mass {
		res = e
	}
	return res
}
