package piscine

func Split(str, charset string) []string {
	countS := 0
	countC := 0
	countE := 0

	for range str {
		countS++
	}
	for range charset {
		countC++
	}

	for i := 0; i < countS-countC; i++ {
		if str[i:i+countC] == charset {
			countE++
		}
	}

	result := make([]string, countE+1)
	m := 0
	n := 0

	for i := 0; i <= countS-countC; i++ {
		if str[i:i+countC] == charset {
			result[m] = string(str[n:i])
			m++
			n = countC + i
		}
		if i == countS-countC {
			result[m] = string(str[n:])
		}
	}
	return result
}
