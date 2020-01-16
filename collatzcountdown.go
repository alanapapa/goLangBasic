package piscine

func CollatzCountdown(start int) int {
	res := 1
	if start <= 0 {
		return -1
	} else if start > 0 {
		for start != 1 {
			if start%2 == 0 {
				start = start / 2
				res++
			} else {
				start = 3*start + 1
				res++
			}
		}
	}
	return res
}
