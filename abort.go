package piscine

func Abort(a, b, c, d, e int) int {
	var res int
	mas := []int{a, b, c, d, e}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if mas[j] < mas[i] {
				mas[i], mas[j] = mas[j], mas[i]
			}
		}
	}

	res = mas[2]
	return res
}
