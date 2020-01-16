package piscine

func ActiveBits(n int) uint {
	var count int
	flag := false
	if n < 0 {
		n = -n
		flag = true
	}
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	if flag {
		count++
	}
	result := uint(count)
	return result
}
